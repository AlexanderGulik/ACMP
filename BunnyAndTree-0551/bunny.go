package main
import (
  "fmt"
  "math"
)

func main(){
  var R, r, h, b float64
  fmt.Scan(&R)
  fmt.Scan(&r)
  fmt.Scan(&h)
  fmt.Scan(&b)
  //S := 2*math.Pi*(r+h)
  //SR := math.Pi * R * R
  if b+r >= R{
    fmt.Println("YES")
    return
  }
  result := math.Sqrt(R*R - r*r) + (b - r)
  if result >= h{
    fmt.Println("YES")
  } else {
    fmt.Println("NO")
  }

}
