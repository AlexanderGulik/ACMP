


package main
import(
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n)
  step := make([]int, n+3)
  for i := 1; i <= n; i++{
    fmt.Scan(&step[i])
  }
  var result int
  curr := 1
  for i := 1; i <=10; i++{
    currStep := step[curr]
     if currStep == 10 {
      result += 10 + step[curr+1] + step[curr+2]
      curr++
     } else {
       d := step[curr+1]
       if currStep + d == 10 {
         result += 10 + step[curr+2]
       } else {
          result += currStep + d
       }
       curr +=2
    }
  }
  fmt.Println(result)
}
