package main
import (
  "fmt"
  "math"
)

func main(){
  var count int
  fmt.Scan(&count)
  count = 3*int(math.Pow(2,float64(count-1)))
  fmt.Println(count)
  return 
}

