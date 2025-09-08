package main

import (
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n)
  wins := make([]int, n)
  for i := 0; i < n; i++{
    fmt.Scan(&wins[i])
  }
  dp := make([][]int, n)
  for i := range dp {
    dp[i] = make([]int, n)
  }
  for i := 0; i < n ; i++{
    j := n - 1 - i
    dp[i][j] = wins[i]
  }
  for i := n-1; i >=0; i--{
    for j := n-1; j >=0; j--{
      if i+j >= n-1 {
        continue
      }
      var moveRight, moveUp int
      if j+1 < n && i+j+1 <= n-1 {
        moveRight = dp[i][j+1]
      }
      if i+1 < n && i+j+1 <= n-1 {
        moveUp = dp[i+1][j]
      }
      if (i+j) % 2 == 0{
        dp[i][j] = max(moveUp, moveRight)
      } else {
        dp[i][j] = min(moveUp, moveRight)
      }
    }
  }
    fmt.Println(dp[0][0])
}
func min(a, b int) int {
  if a < b {
    return a
  } 
  return b
}

func max(a, b int) int {
  if a > b {
    return a
  } 
  return b
}

