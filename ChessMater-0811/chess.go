package main
import (
  "fmt"
)

func main (){
  var n, m int
  fmt.Scan(&n, &m)
  boardCurr := make([][]rune, n)
  for i, _ := range boardCurr {
    boardCurr[i] = make([]rune, m)
    var row string
    fmt.Scan(&row)
    for j := 0; j < m; j++{
      boardCurr[i][j] = rune(row[j])
    }
  }

  pattern1 := createPattern(n, m, 'W')
  pattern2 := createPattern(n, m, 'B')
  change1 := countChange(boardCurr, pattern1, n, m)
  change2 := countChange(boardCurr, pattern2, n, m)

  var result [][2]int
  if len(change1) <= len(change2) {
    result = change1
  } else {
    result = change2
  }
  fmt.Println(len(result))
  for i := 0; i < len(result); i++{
    fmt.Println(result[i][0], result[i][1])
  }
}

func createPattern(n, m int, startChar rune) [][]rune{
  board := make([][]rune, n)
  for i := range board{
    board[i] = make([]rune, m)
  }
  for i := 0; i < n; i++{
    currColor := startChar
      if i %2 != 0 {
        if startChar == 'W' {
          currColor = 'B'
        } else {
          currColor = 'W'
        }
      }

    for j :=0; j < m; j++{
      board[i][j] = currColor
       if currColor == 'W' {
          currColor = 'B'
        } else {
          currColor = 'W'
        }
    }
  }
  return board
}

func countChange(currBoard, trueBoard [][]rune, n, m int) [][2]int{
  result := make([][2]int, 0)
  for i := 0; i < n; i++{
    for j := 0; j < m; j++{
      if currBoard[i][j] != trueBoard[i][j] {
        result = append(result, [2]int{i+1,j+1})
      }
    }
  }
  return result
}
