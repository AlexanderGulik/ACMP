package main

import (
  "fmt"
  "sort"
  "strconv"
)

func main() {
  var n, k int
  fmt.Scan(&n, &k)
  letter := findPrimeNumber(n)
  sort.Strings(letter)
  ind := binarySearch(letter, k)
  fmt.Println(ind+1)
}

func binarySearch(arr []string, k int) int {
  left := 0
  right := len(arr)-1
  n := strconv.Itoa(k)
  for left <= right{
    mid := left + (right-left)/2
    if arr[mid] == n{
      return mid
    }
    if arr[mid] < n {
      left = mid+1
    } else {
    right = mid -1
   }
  }
return 0
}

func findPrimeNumber(n int) []string {
  result := []string{}
  for i := 1; i <= n;i++{
      result = append(result, strconv.Itoa(i))
  }
  return result
}


