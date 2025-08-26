package main

import (
    "fmt"

)

func main() {
    var n, m int
    fmt.Scan(&n, &m)
    result := 2 % m
    for i := 0; i <n; i++{
      result = (result*result) %m
    }
    fmt.Println(result)
  
}
