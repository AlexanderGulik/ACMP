package main
import (
  "fmt"
)

func main() {
  var n int
  result := 0
  fmt.Scan(&n)
  a1 := n/4
  if n < 4 {
    fmt.Println(0)
    return
  }
  for a := 1; a <= a1; a++{
    for b := a; b <= (n-a)/3; b++{
      result = result + (n-a-b)/2-b+1
    }
  }
  fmt.Println(result)
}
