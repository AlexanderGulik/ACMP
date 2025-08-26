package main 
import (
  "fmt"
  "math"
)

func main(){
  var n int
  fmt.Scan(&n)
  result := countPrimaryNumber(n)
  fmt.Println(result)

}

func countPrimaryNumber(n int) int {
  result := 0
  for i := n+1; i < 2*n; i++{
    if isPrime(i) {
      result++
    }
  }
  return result
}

func isPrime(n int) bool {
  if n == 2 || n == 1 {
    return true
  }
  if n % 2 == 0 {
    return false
  }
  sqrtN := int(math.Sqrt(float64(n)))
  for i := 3; i<= sqrtN; i+=2{
    if n % i == 0 {
      return false
    }
  }
  return true
}
