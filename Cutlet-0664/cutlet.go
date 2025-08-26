package main
import (
  "fmt"
)

func main (){
  var k, m, n int
  fmt.Scan(&k, &m, &n)
  result := 0
  if n <= k{
    result = 2*m
  } else {
    if n*2 % k == 0{
      result = (n*2/k)*m
    } else {
      result = (n*2/k+1)*m
    }
  }
    fmt.Println(result)
}
