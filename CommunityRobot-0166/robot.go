package main

import (
  "fmt"
)

func main() {
  var k, n int
  fmt.Scan(&k, &n)
  robotLife := make([]int64, 3)
  robotLife[0] = int64(k)
  for i := 1; i <n; i++{
    total := robotLife[0]+robotLife[1]+robotLife[2]
    newRobots := maxProduct(total)
    robotLife[2] = robotLife[1]
    robotLife[1] = robotLife[0]
    robotLife[0] = newRobots
  }
  result := robotLife[0]+ robotLife[1]+ robotLife[2]
  fmt.Println(result)
}

func maxProduct(n int64) int64 {
  groups5 := n /5
  remaing := n %5
  for remaing <= n {
    if remaing % 3 == 0 {
      return groups5*9 + (remaing/3) * 5
    }
    groups5--
    remaing += 5
    if groups5 < 0 {
      break
    }
  }
  n = n/3
  return n*5
}
