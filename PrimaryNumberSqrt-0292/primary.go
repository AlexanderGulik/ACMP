package main

import (
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n) 
  result := PCK(n)
  fmt.Println(result)
}
func PCK(n int) int {
  if isPrimary(n) {
    return n
  }
  if n < 10{
    return 0
  }
  sum := sumDigit(n)
  return PCK(sum)
}
func sumDigit(n int)int{
  sum := 0
  for n > 0 {
    sum += n % 10
    n /=10
  }
  return sum
}
func isPrimary(n int)bool{
  if n == 1 {
    return false
  }
  if n == 2{
    return true
  }
  for i := 2; i*i <= n; i++ {
    if n % i == 0 {
      return false
    }
  }
  return true
  
}
