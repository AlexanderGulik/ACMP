package main

import (
  "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
    fmt.Scan(&nums[i])
  }

  for {
    result := subseq(nums)
    if len(result) == len(nums){
      break
    }
    nums = result
  }
  fmt.Println(len(nums))
  for i := 0; i < len(nums); i++{
    fmt.Printf("%d ", nums[i])
  }
}

func subseq(nums []int) []int {
  if len(nums) <= 2 {
    return nums
  }

  result := []int{}
  result = append(result, nums[0])
  for i := 1; i < len(nums)-1; i++ {
    leftGreater := false
    rightGreater := false
    for j := 0; j < i; j++{
      if nums[j] > nums[i] {
        leftGreater = true
        break
      }
    }
    
    for j := i + 1; j < len(nums); j++{
      if nums[j] > nums[i] {
        rightGreater = true
        break
      }
    }

    if leftGreater && rightGreater {
      continue
    }
    result = append(result, nums[i])
  
  }
  result = append(result, nums[len(nums)-1])
  return result
}
