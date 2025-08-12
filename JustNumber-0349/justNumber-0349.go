package main

import (
  "os"
  "strings"
  "fmt"
  "strconv"
  "bufio"
  "math"
)
func main(){
  content, _ := os.ReadFile("input.txt")
  nums := strings.Split(strings.TrimSpace(string(content)), " ")
  result := []int{}
  start, _ := strconv.Atoi(nums[0])
  end, _ := strconv.Atoi(nums[1])
  primes :=  primeNumber(end)

  for _, p := range primes{
    if p >= start {
      result = append(result,p)
    }
  }
  outFile, _ := os.Create("output.txt")
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)
  if len(result) == 0 {
    fmt.Fprintln(writer, "Absent")
  } else {
  for _, line := range result {
    fmt.Fprintln(writer, line)
   }
  }
  writer.Flush()
}

func primeNumber(num int) []int{
  if num < 2 {
    return []int{}
  }
  sieve := make([]bool, num+1)
  for i := 2; i <= num; i++ {
    sieve[i] = true
  }
  for i := 2; i <= int(math.Sqrt(float64(num)));i++{
    if sieve[i] {
      for j := i * i; j <= num; j+=i {
        sieve[j] = false
      }
    }
  }
  primes := []int{}
  for i, isPrime := range sieve {
    if isPrime {
      primes = append(primes, i)
    }
  }
  return primes
}


/*func justNumber(num int) bool {
  if num == 2 {
    return true
  }
  if num < 2 {
    return false
  }
  if num % 2 == 0 {
    return false
  }
  sqrtNum := int(math.Sqrt(float64(num)))
  for i := 3; i <= sqrtNum; i+=2{
    if num % i == 0 {
      return false
    }
  }
  return true
}*/


