package main
import (
  "fmt"
  "math"
)

func main(){
  var k, n, m, d int
  fmt.Scan(&k, &n, &m, &d)
  den := 1- (1/float64(k) + 1/float64(m) + 1/float64(n))
  if den <= 0 {
    fmt.Println(-1)
    return
  } 
  x := float64(d) / den
  if math.Abs(x-math.Round(x)) < 1e-9 {
    xInt := int(math.Round(x))
    if xInt >0 && xInt%k==0&& xInt%n==0 && xInt%m ==0 {
      fmt.Println(xInt)
      return
    }
  }
  fmt.Println(-1)
}

