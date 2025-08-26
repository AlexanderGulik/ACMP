package main

import(
    "fmt"
)

func main() {
  var n int
  fmt.Scan(&n)
  stats := make([][]int, n)
  for i := 0;i < n; i++{
    stats[i] = make([]int, 3)
  }
  for i :=0; i< n; i++{
    for j :=0; j< 3; j++{
      fmt.Scan(&stats[i][j])
    }
  }
  var time int
  fmt.Scan(&time)
  var  current int
  for t := 0; t < time; t++{
    sumV := 0
    for i := 0; i < n; i++{
      if stats[i][0] <= t && t < stats[i][1]{
        sumV += stats[i][2]
       }
      }
      current += sumV
      if current < 0 {
        current = 0
      }
    }
  
  fmt.Println(current)
}
