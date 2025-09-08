package main

import (
  "fmt"
)

func main(){
  var n int
  var p float64
  fmt.Scan(&n)
  fmt.Scan(&p)
  min, max := 30.0, 4000.0
  for i := 1; i < n; i++{
    var t float64 
    var s string
    fmt.Scan(&t, &s) 
    x := (p+t)/2
    if s == "closer" {
      if t > p && min < x {
        min = x
      }
      if t < p && max > x { 
        max = x
      }
    } else {
      if t < p && min < x {
        min = x
      }
      if t > p && max > x {
        max = x
      }
    }
    p = t
  }
  fmt.Printf("%.6f %.6f",min, max)
}

