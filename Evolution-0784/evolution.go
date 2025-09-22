package main
import (
  "fmt"
)

func main() {
  var n, a, b int64
  fmt.Scan(&n, &a, &b)

  for a != b {
    if a > b {
      a = a / 2
    } else {
      b = b / 2
    }
  }

  fmt.Println(a)

}
