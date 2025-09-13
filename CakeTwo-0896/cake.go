package main

import (
  "fmt"
)

func main() {
  var n, m int
  fmt.Scan(&n, &m)
  result := n+m - abs(n, m)
 
  fmt.Println(result)

}

func abs(n, m int) int {
  for m != 0 {
    n, m = m, n%m
  }
  return n
}
