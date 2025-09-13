package main

import (
  "fmt"
  "sort"
)

func main() {
  var n int
  fmt.Scan(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
    fmt.Scan(&nums[i])
  }
  var sum float64
  fmt.Scan(&sum)
  sort.Ints(nums)
  for i := 0; i < n; i++{
    curr := (sum + float64(nums[i]))/2.0 
    if curr > sum {
      sum = curr
    }
  }
  fmt.Printf("%.6f", sum)
}
