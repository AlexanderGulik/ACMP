package main
import (
  "fmt"
)

func main(){
  var m, n int 
  fmt.Scan(&m, &n)
  a := make([][]int, m)
  for i := range a {
    a[i] = make([]int, n)
  }
  for i := 0; i < m; i++{
    for j := 0; j < n; j++{
      fmt.Scan(&a[i][j])
    }
  }
  var DownInt int
  currDownInt := make([]int, m)
  for i:=0; i < m; i++{
    currDownInt[i] = min(a[i])
  }

  DownInt = max(currDownInt)
  columnMax := make([]int, n)
  for j := 0; j < n; j++{
    column := getColumn(a, j)
    columnMax[j] = max(column)
  }
  UpInt := min(columnMax)
  fmt.Println(DownInt, UpInt)
}

func getColumn(matrix [][]int, colIdx int) []int {
  curr := make([]int, len(matrix))
  for i := range matrix{
    curr[i] = matrix[i][colIdx]
  }
  return curr
}

func max(arr []int) int{
  curr := arr[0]
  for i := 0; i < len(arr); i++{
    if arr[i] > curr{
      curr = arr[i]
    }
  }
  return curr
}
func min(arr []int) int{
  curr := arr[0]
  for i := 0; i < len(arr); i++{
    if arr[i] < curr{
      curr = arr[i]
    }
  }
  return curr
}
