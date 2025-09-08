package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  var result int
  if n & (n-1) == 0{
    result = 0
  } else {
    m := 1
    for m <= n {
      m <<=1
    }
    m >>= 1
    result = n-m
  }
  fmt.Println(result)
}
