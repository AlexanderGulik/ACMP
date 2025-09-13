package main

import (
  "bufio"
  "fmt"
  "os"
  "strconv"
  "strings"
)

func main() {
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  nm := strings.Fields(scanner.Text())
  n, _ := strconv.Atoi(nm[0])
  m, _ := strconv.Atoi(nm[1])

  grid := make([][]int, n)
  for i := range grid {
    grid[i] = make([]int, m)
    scanner.Scan()
    row := strings.Fields(scanner.Text())
    for j := range row {
      grid[i][j], _ = strconv.Atoi(row[j])
    }
  }
  dp := make([][]int, n)
  for i := range dp {
    dp[i] = make([]int, m)
  }
  dp[0][0] = 1
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      if dp[i][j] == 0 {
        continue 
      }
      step := grid[i][j]
      if step == 0 {
        continue
      }
      if i+step < n{
        dp[i+step][j] += dp[i][j]
      }
      if j+step < m {
        dp[i][j+step] += dp[i][j]
      }
    }
  }
  fmt.Println(dp[n-1][m-1])
}
