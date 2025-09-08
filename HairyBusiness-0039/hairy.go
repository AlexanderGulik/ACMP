package main
import (
  "fmt"
)
func main() {
  var n int
  fmt.Scan(&n)
  days := make([]int, n)
  for i := 0; i<n; i++{
    fmt.Scan(&days[i])
  }
  lenghtHair := 1
  result := 0
  for i := 0; i < n; i++{
    if IsMax(days[i+1:], days[i]){
      result += lenghtHair * days[i]
      lenghtHair = 0
    } 
    lenghtHair++
  }
  //288 + 92 380
  fmt.Println(result)

}
func IsMax(days []int, n int) bool{
  for i := 0 ; i < len(days); i++{
    if days[i] >= n{
      return false
    }
  }
  return true
}
