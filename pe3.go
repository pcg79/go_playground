package main

import "fmt"
import "os"
import "strconv"

func main() {
  prime_factors := []int{1}

  remainder, _ := strconv.Atoi(os.Args[1])

  for remainder > 1 {
    for v := 2; v < 100000; v++ {
      if remainder % v == 0 {
        remainder = remainder / v
        prime_factors = append(prime_factors, v)
      }
    }
  }

  fmt.Println("Answer is:", prime_factors)
}
