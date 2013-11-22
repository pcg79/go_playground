package main

import "fmt"

func main() {
  var sum = 0
  var fib_prev, fib = 1, 1

  for fib < 4000000 {
    if fib % 2 == 0 {
      sum = sum + fib
    }

    tmp := fib
    fib = fib + fib_prev
    fib_prev = tmp
  }

  fmt.Println("Answer is:", sum)
}
