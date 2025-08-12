package main

import (
  "fmt"
)

func main(){
  var n, k int
  fmt.Scan(&n)
  fmt.Scan(&k)
  result := 0
  for i := 0; i <= k; i++ {
    result = result + (n-2) * i +1
  }
  fmt.Println(result)

}
