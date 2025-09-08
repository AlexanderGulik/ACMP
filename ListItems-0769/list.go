package main

import (
  "fmt"
  "sort"
)

func main(){
  var n int
  fmt.Scan(&n)
  words := make([]string, n)
  for i := 0; i < n; i++{
    var item string
    fmt.Scan(&item)
    words[i] = item
  }
  sort.Strings(words)
  result := 0
  for i := 0; i<n; i++{
    if i < n-1 && isPrefix(words[i], words[i+1]){
      continue
    }
    result++
  }
  fmt.Println(result)
}

func isPrefix(prefix, word string) bool {
  if len(prefix) > len(word) {
    return false
  }
  return word[:len(prefix)] == prefix
}
