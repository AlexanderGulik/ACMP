package main
import (
  "fmt"
)
func gcd(a, b int) int {
  for (b != 0) {
    t := a % b
    a = b
    b = t
  }
  return a
}
func divCount(n int) int {
  if n == 1 {
    return 0
  }
  count := 1
  d := 2
  for d *d <= n {
    if n % d == 0 {
      count++
      n /= d
    } else {
      d++
    }
  }
  return count
}

func main(){
  var n, m int
  fmt.Scan(&n, &m)
  g := gcd(n, m)
  result := divCount(m /g ) + divCount(n/ g)
  fmt.Println(result)
}

