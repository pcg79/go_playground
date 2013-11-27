// A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

// Find the largest palindrome made from the product of two 3-digit numbers.

// Two ways to do this.

// 1) Start w/ two numbers - 999 and 999.  Work down finding numbers that are palindromes.  Cons: that's a lot of multiplying.

// 2) Start at the max number made by multiplying 2 3 digit numbers (999*999 = 998001) and work down to min number made by multiplying 2 3 digit numbers (100*100 = 10000).  Stop when you get a palindrome, then see if it can be made by multiplying 2 3 digit numbers.

// 998001 down to 10000 do |x|

// if x is a palindrome
//   999 down to 100 do |i|
//     return "winner is x" if x mod i is a three digit number.


// is palindrome (x)
//   str = x to str

//   first part = str[0..2]
//   second part = str[3..5]
//   second part = flip second part around.

//   return true if first part = second part

package main

import "fmt"
import "strconv"

func main() {
  for i := (999 * 999); i > (100 * 100); i-- {
    if isPalindrome(i) {
      for j := 100; j < 1000; j++ {
        if i % j == 0 && i / j > 99 && i / j < 1000 {
          fmt.Println("Answer is: ", i)
          fmt.Println("Factors: ", j, i / j)
          return
        }
      }
    }

  }
}

func isPalindrome(x int) bool {
  str := strconv.Itoa(x)

  length := len(str) / 2
  start := str[:length]
  end := str[length:]

  // if start and end are different lengths, that means it's an odd-lengthed number.
  // so we just strip the first number off of "end" since it's the middle digit and
  // doesn't need a mirror.
  if len(start) != len(end) {
    end = str[:len(end) - 1]
  }

  end = Reverse(end)

  return start == end
}

// Taken from http://stackoverflow.com/questions/1752414/how-to-reverse-a-string-in-go
func Reverse(s string) string {
    runes := []rune(s)
    for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
        runes[i], runes[j] = runes[j], runes[i]
    }
    return string(runes)
}