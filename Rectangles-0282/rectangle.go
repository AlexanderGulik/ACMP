package main

import (
  "fmt"
)

func main() {
  var w, h int
  fmt.Scan(&w, &h)
  result := (w*(w+1) * h * (h+1))/4
  fmt.Println(result)
}

