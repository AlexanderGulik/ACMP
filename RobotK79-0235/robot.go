package main

import (
  "fmt"
)

func main(){
  var s string
  fmt.Scan(&s)
  var a[104][104]byte
  dx := [4]int{0, 1, 0, -1}
  dy := [4]int{1, 0, -1, 0}
  m:= 0
  napr := 0
  x:= 51
  y := 51
  a[x][y] = 1
  for i := 0; i < len(s); i++{
  if s[i] == 'R' {
    napr = (napr+5) %4
  }
  if s[i] == 'L' {
    napr = (napr+3) %4
  }
  if s[i] == 'S' {
    x += dx[napr]
    y += dy[napr]
    if a[x][y] == 1 {
      fmt.Println(m+1)
      return
    } else {
      a[x][y] =1
      m++
    }
  }
  }
  fmt.Println(-1)
}
