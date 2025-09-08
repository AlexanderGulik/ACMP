package main

import (
  "fmt"
  "math"
)

func main() {
  var x1, y1 float64
  var x2, y2 float64
  var r, s float64
  fmt.Scan(&x1, &y1)
  fmt.Scan(&x2, &y2)
  fmt.Scan(&r, &s)
  s1 := math.Pi * r*r
  s2 := math.Pi * r*r
  result := 0.0
  d := math.Sqrt(math.Pow((x2-x1),2) + math.Pow((y2-y1),2))
  if d >= 2*r {
    result = s1+s2
  } else if d <= 0 {
    result = s1
  } else {
    angle1 := 2*math.Acos((d*d)/(2*r*d))
    angle2 := 2*math.Acos((d*d)/(2*r*d))
    segment1 := 0.5 * r * r * (angle1 - math.Sin(angle1))
    segment2 := 0.5 * r * r * (angle2 - math.Sin(angle2))
    inter := segment1 + segment2
    result = s1+s2-inter
  }

  if result > s {
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }

}




