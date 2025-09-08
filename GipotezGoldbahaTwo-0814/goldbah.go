package main
import (
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n)
  result := SearchTwoNum(n)
  fmt.Println(result)
}

func SearchTwoNum(n int) int {
  count := 0
  used := make(map[int]bool)
  for i := 2; i <= n; i++{
    if isPrimary(i){
      j := n-i 
      if j >= 2 && isPrimary(j) {
        if !used[i] && !used[j]{
          used[i] = true
          used[j] = true
          count++
        }
      }
    }
  }
  return count
}

func isPrimary(n int) bool {
  if n == 1 {
    return false
  }
  if n == 2 {
    return true
  }
  for i := 2; i*i <= n; i++{
   if n % i == 0 {
     return false
   }
  }
  return true
}
