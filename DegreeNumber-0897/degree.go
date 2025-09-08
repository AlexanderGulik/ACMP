package main
import (
  "fmt"
  "math"
)

func main(){
  var n int
  fmt.Scan(&n)
  for i := 0; i < n; i++{
    var num int
    fmt.Scan(&num)
    if isDegree(num) {
      fmt.Println("YES")
    } else {
      fmt.Println("NO")
    }
  }
    
}

func isDegree(n int) bool {
  if n <= 1 {
    return true
  } 
  max := int(math.Sqrt(float64(n))) + 1
  for i := 2; i <= max; i++{
      p := i*i
      for p <=n {
        if p == n{
          return true
        }
        p *= i
      }
  }
  return false
}
