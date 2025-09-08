package main
import (
  "fmt"
)

func main() {
  var n int64
  fmt.Scan(&n)
  if n % 2 == 1{
    fmt.Println(2, n+2)
  } else {
    fmt.Println(0, 0)
  }
}
