package main
import(
  "fmt"
  "math/big"
)

func main(){
    var n int
    fmt.Scan(&n)
    start := big.NewInt(1) 
    ten := big.NewInt(10) 
    one := big.NewInt(1) 
    two := big.NewInt(2) 
    start.Exp(ten, big.NewInt(int64(n-1)), nil)

    last := new(big.Int).Exp(ten, big.NewInt(int64(n)), nil)
    last.Sub(last, one)
    count := new(big.Int).Sub(last, start)
    count.Add(count, one)
    num := new(big.Int).Add(start, last)
    num.Mul(num, count)
    result := new(big.Int).Div(num, two)
    fmt.Println(result)
}

// сука почему я так поздно заметил что можно было решить это через строки, не используя биг инт
