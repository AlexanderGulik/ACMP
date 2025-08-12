package main
import (
  "sort"
  "math"
  "fmt"
)

func main() {
  var count int
  fmt.Scan(&count)
  s := make([][2]int, count)
  for i := 0; i<count;i++{
    fmt.Scan(&s[i][0], &s[i][1])
  }
  a:= make(map[float64]bool)
  for i := 0; i < count; i++{ 
    for j := i+1; j < count; j++{
      dist := solve(s[i][0],s[i][1], s[j][0],s[j][1])
      a[dist] = true
    }
  }
  distance := make([]float64, 0, len(a))
  for dist := range a {
    distance = append(distance, dist)
  }
  sort.Float64s(distance)
  fmt.Println(len(distance))
  for _, dist := range distance {
    fmt.Printf("%.9f\n", dist)
  }
}

func solve(x1, y1, x2, y2 int) float64 {
  dx := float64(x1-x2)
  dy := float64(y1-y2)
  return math.Sqrt(dx*dx+dy*dy)
}
