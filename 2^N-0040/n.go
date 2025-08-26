package main
import (
  "fmt"
  "math/big"
)

func main(){
  var n int
  fmt.Scan(&n)
  result := new(big.Int).Lsh(big.NewInt(1), uint(n))
  fmt.Println(result)
}
