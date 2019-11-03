-- Comments
{-
    Haskell uses type inference and is statically typed 
    Int -2^63 to 2^63
    Integer: unbounded whole number
    Float: single precision
    Double: up to 11 decimal point numbers
    Bool
    Char: single unicode
    Tuple: lists made up of different data types
-}

import Data.List
import Data.Typeable
import System.IO

maxInt = maxBound :: Int
minInt = minBound :: Int
sumOfNums = sum [1..1000]

addEx = 5 + 4
subEx = 5 + 4
multEx = 5 + 4
divEx = 5 + 4

negNumEx = 5 + (-4)

modEx = mod 5 4
modEx2 = 5 `mod` 4

num9  = 9 :: Int
sqrtOf9 = sqrt(fromIntegral num9)

piVal = pi
ePow9 = exp 9
log9 = log 9
squared9 = 9 ** 2
truncateVal = truncate 9.999
roundVal = round 9.999
ceilVal = ceiling 9.999
floorVal = floor 9.999

-- also:
--  sin,   cos,   tan
-- asin,  acos,  atan
--  sinh,  cosh,  tanh
-- asinh, acosh, atanh

trueAndFalse = True && False
trueOrFalse = True || False
notTrue = not(True)

primeNumbers = [3,5,7,11]

morePrimes = primeNumbers ++ [13,17,19.23,29]

favNums = 2 : 7 : 21 : 66 :[]

multList = [[3,5,7],[11,13,17]]
morePrimes2 = 2 : morePrimes
lenPrime = length morePrimes2
revPrime = reverse morePrimes2
isListEmpty = null morePrimes2
secondPrime = morePrimes2 !! 1
firstPrime = head morePrimes2
lastPrime = last morePrimes2
primeInit = init morePrimes2
first3Primes = take 3 morePrimes2
removedPrimes = drop 3 morePrimes2
is7InList = 7 `elem` morePrimes2
maxPrime = maximum morePrimes2
minPrime = minimum morePrimes2
sumOfPrimes = sum morePrimes2

newList = [2,3,5]
prodPrimes = product newList
zeroToTen = [0..10]
evenList = [2,4..20]
letterList = ['A', 'C'..'Z']
infinPow10 = [10,20..] -- an infinite list, lazy evaluated
many2s = take 10 (repeat 2)
many3s = replicate 10 3
cycleList = take 10 (cycle [1,2,3,4,5])
listTimes2 = [x * 3 | x <- [1..10], x * 3 <= 50]
divisBy9N13 = [x | x <- [1..500], x `mod` 13 == 0, x `mod` 9 == 0] -- OR filters
sortedList = sort [9, 1, 8, 3, 4, 7, 6]
sumOfLists = zipWith (+) [1,2,3,4,5] [6,7,8,9,10]

-- use ":t sqrt" in ghci

-- main = putStrLn  "hello, world"
main = do
    print sumOfLists
    putStrLn("Please enter your name: ")
    name <- getLine
    putStrLn("Hi, " ++ name ++ "!")
