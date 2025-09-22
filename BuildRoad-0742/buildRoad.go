package main

import (
	"fmt"
	"math/big"
)

func main() {
	var n int
	fmt.Scan(&n)
	x := n * (n - 1) / 2
	result := new(big.Int).Exp(big.NewInt(3), big.NewInt(int64(x)), nil)
	fmt.Println(result)

}
