package main

import (
  "fmt"
  "strings"
  "strconv"
  "os"
)

func main() {
  nums, _ := os.ReadFile("input.txt")
  data := strings.TrimSpace(string(nums))
  parts := strings.Fields(data)
  var result string
  for _, char := range parts {
    num, _ := strconv.Atoi(char)
    if numberSmit(num) {
      result += "1"
    } else {
      result += "0"
    }
  }
  fmt.Println(result)
}

func numberSmit(n int)bool{
  if isPrime(n) {
    return false
  }
  digitSum := sumDigit(n)
  primeSum := sumPrime(n)
  return digitSum == primeSum
}

func isPrime(n int) bool {
  if n < 2 {
    return false
  }
  for i := 2; i*i <= n; i++{
    if n % i ==0 {
      return false
    }
  }
  return true
}

func sumDigit(n int) int {
  sum := 0
  for n > 0 {
    sum += n%10
    n /= 10
  }
  return sum
}

func sumPrime(n int) int {
  sum := 0
  temp := n

  for i :=2; i <=temp; i++{
    for temp %i == 0{
      sum += sumDigit(i)
      temp /= i
    }
  }
  return sum
}
