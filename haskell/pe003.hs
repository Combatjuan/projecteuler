-- If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
--
-- Find the sum of all the multiples of 3 or 5 below 1000.
-- sum [x for x in 1..999 where (mod x 3 == 0 || mod x 5 == 0)]

import qualified Data.Numbers.Primes as Primes

main = do
	print (sum $ take 10 Primes.primes)

