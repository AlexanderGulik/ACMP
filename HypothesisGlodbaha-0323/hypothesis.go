package main
import (
  "fmt"
  "math"
)

func main(){
  var n int
  fmt.Scan(&n)
  n1, n2 := findTwoPrimeNumber(n)
  fmt.Printf("%d %d", n1, n2)
}

func findTwoPrimeNumber(num int) (int, int){
  for i := 2; i < num; i += 1{
      if isPrime(i) && isPrime(num - i) {
        return i, num-i
    }
  }
  return 0, 0
}

func isPrime(n int) bool{
  
  num := int(math.Sqrt(float64(n)))
  for i := 2; i <=num; i+=1 {
   if n % i == 0 {
      return false
    }
  }
  return true
}
