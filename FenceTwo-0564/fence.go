package main
import (
  "fmt"
  "math"
)

func main() {
 var n int
  fmt.Scan(&n)
  nums := make([]float64, n)
  for i := 0; i < n; i++{
    fmt.Scan(&nums[i])
  }
  var result float64
  index := [3]int{0,0,0}
  for i := 0; i < n; i++ {
    for j := i+1; j<n; j++{
      for k := j+1; k < n; k++{
       p := (nums[i] + nums[j] + nums[k])/2
       s := math.Sqrt(p*(p-nums[i]) *(p-nums[j])*(p-nums[k]))
       if s > result {
        result = s
        index[0] = i
        index[1] = j
        index[2] = k
       }
      }
    }
  }
  if result == 0 {
    fmt.Println("-1")
    return
  }
  fmt.Println(result)
  fmt.Println(index[0]+1, index[1]+1, index[2]+1)
}
