package main

import (
  "fmt"
)

func main() {
  fmt.Println(square(3))
  fmt.Println(add(2,7))
  fmt.Println(substract(9,3))
}

 func square(number int) int {
  return number * number
 }
  
 func add(a float64, b float64) (sum float64) {
  return a + b
 }

 func substract(a,b float64) (difference float64) {
  difference = a - b
  return
 }