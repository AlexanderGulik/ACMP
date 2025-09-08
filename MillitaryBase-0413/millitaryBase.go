package main

import (
  "fmt"
)

func main() {
  var n, m int
  fmt.Scan(&n, &m)
  plase := make([][]rune, n)
  visited := make([][]bool, n)
  for i := 0; i < n; i++{
    plase[i] = make([]rune, m)
    visited[i] = make([]bool, m)
    var row string
    fmt.Scan(&row)
    for j, char := range row{
      plase[i][j] = char
    }
  }
    count := 0
  direct := [][]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
  
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      if plase[i][j] == '#' && !visited[i][j] {
        count++
        queue := [][]int{{i,j}}
        visited[i][j] = true
        for len(queue) > 0 {
          current := queue[0]
          queue = queue[1:]
          for _, dir := range direct {
            ni, nj := current[0]+dir[0], current[1]+dir[1]
              if ni >= 0 && ni < n && nj >= 0 && nj <= m && plase[ni][nj] == '#' && !visited[ni][nj]{
                visited[ni][nj] = true
                queue = append(queue, []int{ni, nj})
              }
          }
        }
      }
    }
  }
  fmt.Println(count)
}
