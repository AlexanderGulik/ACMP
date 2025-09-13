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
  sort.Ints(nums)
  minRatio := float64(1<<63-1)
  bestStart := 0
  bestEnd := 0
   for i := 0; i < n-1; i++{
     ratio2 := float64(nums[i+1] - nums[i])
     if ratio2 < minRatio {
      minRatio = ratio2
      bestStart = i
      bestEnd = i+1
     }
     if i < n-2 {
      ratio3 := float64(nums[i+2]- nums[i]) /2.0
      if ratio3 < minRatio {
        minRatio = ratio3
        bestStart = i
        bestEnd = i+2
      }
     }
   }
     fmt.Println(nums[bestStart], nums[bestEnd])
  
  
}
