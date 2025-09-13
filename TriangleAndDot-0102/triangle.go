package main

import (
  "fmt"
  "math"
)

func main() {
 var x1, y1, x2, y2, x3, y3 float64
 var x, y float64
 fmt.Scan(&x1, &y1, &x2, &y2, &x3, &y3)
 fmt.Scan(&x, &y)
  areaPAB := triangleArea(x, y, x1, y1, x2, y2)
  areaPBC := triangleArea(x, y, x2, y2, x3, y3)
  areaPCA := triangleArea(x, y, x3, y3, x1, y1)
  areaABC := triangleArea(x1, y1, x2, y2, x3, y3)
  if math.Abs(areaABC - (areaPAB + areaPCA + areaPBC)) < 1e-9{
    fmt.Println("In")
  } else {
    fmt.Println("Out")
  }
}

func triangleArea(x1, y1, x2, y2, x3, y3 float64) float64 {
  return math.Abs((x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0)
}
