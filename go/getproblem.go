package main

import (
	"flag"
	"fmt"
	"text/template"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func getProblem(problem int) (string, string, error) {
	url := fmt.Sprintf("https://projecteuler.net/problem=%d", problem)
	response, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()

	z := html.NewTokenizer(response.Body)

	var descriptionLevel int
	var more bool
	parsing := true
	inDescription := false
	level := 0
	description := make([]string, 0)
	var document string
	var raw []string
	for parsing {
		token := z.Next()
		raw = append(raw, string(z.Raw()))
		switch token {
		case html.ErrorToken:
			parsing = false
		case html.TextToken:
			if inDescription {
				description = append(description, string(z.Text()))
			}
		case html.StartTagToken:
			tagName, _ := z.TagName()
			if string(tagName) == "div" {
				var key []byte
				var val []byte
				more = true
				for more {
					key, val, more = z.TagAttr()
					if string(key) == "class" && string(val) == "problem_content" {
						inDescription = true
						descriptionLevel = level
					}
				}
			} else if inDescription && string(tagName) == "p" {
				description = append(description, "\n")
			}
			level++
		case html.EndTagToken:
			level--
			if level <= descriptionLevel {
				inDescription = false
			}
		}
	}
	document = strings.Join(raw, "")
	return strings.TrimSpace(strings.Join(description, "")), document, nil
}

type ProblemInfo struct {
	Number int
	Description string
}

func writeProblem(number int, description string, page string) {
	problem := ProblemInfo{number, description}
	const goTemplate string =
`package main
/*
Implements Project Euler Problem {{.Number}}:
https://projecteuler.net/problem={{.Number}}
{{.Description}}
*/

import (
	"fmt"
)

func main() {
	fmt.Println("Project Euler Problem {{.Number}}")
}
`
	tplate, err := template.New("GoTemplate").Parse(goTemplate)
	if err != nil {
		panic(err)
	}

	// Write out a template .go file
	goFilename := fmt.Sprintf("p%03d.go", number)
	sourceFile, err := os.OpenFile(goFilename, os.O_WRONLY | os.O_CREATE | os.O_EXCL, os.ModePerm & 0644)
	if os.IsExist(err) {
		fmt.Println(goFilename + " already exists.")
		return
	} else if err != nil {
		panic(err)
	}
	defer sourceFile.Close()
	tplate.Execute(sourceFile, problem)

	// Write out a template .html documentation file.
	docFilename := fmt.Sprintf("p%03d.html", number)
	docFile, err := os.OpenFile(docFilename, os.O_WRONLY | os.O_CREATE | os.O_EXCL, os.ModePerm & 0644)
	if os.IsExist(err) {
		fmt.Println(docFilename + " already exists.")
		return
	} else if err != nil {
		panic(err)
	}
	defer docFile.Close()
	docFile.Write([]byte(page))
}

func main() {
	// Set up argument handling.
	var number = flag.Int("number", 0, "Specify a problem to download.")
	flag.Parse()

	if *number == 0 {
		fmt.Println("Please specify a problem.")
		return
	}
	description, page, err := getProblem(*number)
	if err != nil {
		panic(err)
	}
	writeProblem(*number, description, page)
}

