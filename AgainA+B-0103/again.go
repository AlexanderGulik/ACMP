package main
import (
  "fmt"
  "math/big"
)

func main() {
  var a, b big.Int
  fmt.Scan(&a, &b)
  result := big.NewInt(0)
  result.Add(&a, &b)
  fmt.Println(result)

}


