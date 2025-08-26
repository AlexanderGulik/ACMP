package main
import (
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n)
  for n % 2 == 0 {
    n /= 2
  }
  for n % 5 == 0 {
    n /= 5
  }
  if n == 1 {
    fmt.Println("NO")
  } else {
    fmt.Println("YES")
  }


}
