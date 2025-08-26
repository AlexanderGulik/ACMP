package main
import (
  "fmt"
  "math"
)

func main() {
  var k int
  var t float64
  fmt.Scan(&k, &t)
  nums := make([][3]int, k)
  for i := 0; i < k; i++{
    var x, y, r int
    fmt.Scan(&x, &y, &r)
    nums[i][0] = x
    nums[i][1] = y
    nums[i][2] = r
  }
  currMin := 29000000.0
  for i := 0; i < k; i++{
    x1 := nums[i][0]
    y1 := nums[i][1]
    r1 := nums[i][2]
    for j := i+1; j < k; j++{
    x2 := nums[j][0]
    y2 := nums[j][1]
    r2 := nums[j][2]
    d := math.Sqrt(math.Pow(float64(x2-x1), 2)+math.Pow(float64(y2 - y1), 2) ) - float64(r1+r2)
    d = d/2.0
    if currMin > d {
      currMin = d
    }}
  }
  
  fmt.Printf("%.2f", min(currMin, t))
}

func min(a, b float64) float64 {
  if a < b {
    return a
  }
   return b
}


