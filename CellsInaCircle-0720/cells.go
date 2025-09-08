package main

import (
  "fmt"
  "math"
)

func main () {
  var r float64
  var l int
  fmt.Scan(&r, &l)
  r_scaled := r / float64(l)
  r2 := r_scaled * r_scaled
  count := 0
  max_i := int(math.Floor(r_scaled))
  j_max := max_i
  for i := 0; i <= max_i; i++{
    for j_max >= 0 && float64((i+1)*(i+1) + (j_max+1) * (j_max+1)) > r2{
      j_max--
    }
    if j_max < 0 {
      break
    }
      count += j_max +1
  }
  result := 4 * count
  fmt.Println(result)


}
