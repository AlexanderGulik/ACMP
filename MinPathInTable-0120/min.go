package main
import (
  "fmt"
)

func main () {
  var n, m int
  fmt.Scan(&n, &m)
  matrix := make([][]int, n)
  dp := make([][]int, n)
  for i := 0; i < n; i++{
    dp[i] = make([]int, m)
    matrix[i] = make([]int, m)
  }

  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      fmt.Scan(&matrix[i][j])
    }
  }
  dp[0][0] = matrix[0][0]
  for j := 1; j<m; j++{
    dp[0][j] = dp[0][j-1] + matrix[0][j]
  }

  for i := 1; i<n; i++{
    dp[i][0] = dp[i-1][0] + matrix[i][0]
  }
  
  for i := 1; i < n; i++{
    for j := 1; j < m; j++{
      dp[i][j] = min(dp[i-1][j], dp[i][j-1]) + matrix[i][j]
    }
  }
  fmt.Println(dp[n-1][m-1])
}

func min(a, b int) int {
    if a < b {
      return a
    }
    return b
}
