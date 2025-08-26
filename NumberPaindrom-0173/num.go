package main
import (
  "fmt"
  "strconv"
)

func main(){
  var n int64
  fmt.Scan(&n)
  result := []int{}
  for i := 2; i <=36; i++{
    curr := strconv.FormatInt(int64(n), i)
    if palindorm(curr) {
      result = append(result, i)
    }
  }
  if len(result) == 0 {
    fmt.Println("none")
  } else if len(result) == 1 {
    fmt.Printf("unique \n%d", result[0])
  } else {
    fmt.Printf("multiple\n")
    for i := 0; i < len(result); i++{
      fmt.Printf("%d ", result[i]) 
    }
  }

}

func palindorm(a string) bool {
  left := 0
  right := len(a)-1
  for left < right {
    if a[left] != a[right] {
      return false
    }
    left++
    right--
  }
  return true
}
