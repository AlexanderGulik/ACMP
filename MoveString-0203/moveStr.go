package main

import (
  "fmt"
)

func main(){
  var s1, s2 string
  fmt.Scan(&s1, &s2)
  result := cicleMove(s1, s2)
  fmt.Println(result)
}

func cicleMove(s1, s2 string) int {
  for i := 0; i < len(s1); i++{
    currStr := rotateString(s1, i)
    if currStr == s2 {
      return i
    }
  }
  return -1
}

func rotateString(s string, n int) string {
  if len(s) == 0 {
    return s
  }
  n = n % len(s)
  if n < 0 {
    n+= len(s)
  }
  return s[len(s)-n:] + s[:len(s)-n]
}
