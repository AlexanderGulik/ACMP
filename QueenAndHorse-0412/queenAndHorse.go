package main

import (
  "fmt"
)

func main() {
  var whiteKing, queen, blackKing string
  fmt.Scan(&whiteKing, &queen, &blackKing)
  board := make([][]int, 8)
  for i := 0; i < len(board); i++{
    board[i] = make([]int, 8)
  }
  whiteKingX, whiteKingY := convertChess(whiteKing)
  queenX, queenY := convertChess(queen)
  blackKingX, blackKingY := convertChess(blackKing)
  board[whiteKingX][whiteKingY] = 1
  board[blackKingX][blackKingY] = 3
  if (whiteKingX+1 == blackKingX && whiteKingY == blackKingX) || (whiteKingX-1 == blackKingX && whiteKingY == blackKingX) ||
     (whiteKingX+1 == blackKingX && whiteKingY+1 == blackKingX) || (whiteKingX-1 == blackKingX && whiteKingY-1 == blackKingX) ||
     (whiteKingX-1 == blackKingX && whiteKingY+1 == blackKingX) || (whiteKingX+1 == blackKingX && whiteKingY-1 == blackKingX) {
        fmt.Println("NO")
      return
     }
  board[queenX][queenY] = 2
  if isQueenAttac(queenX, queenY, blackKingX, blackKingY, board) {
    fmt.Println("YES")
  }  else {
    fmt.Println("NO")
  }

}

func isQueenAttac(qx, qy, kx, ky int, board [][]int) bool {
    if qx == kx || qy == ky {
      return isPathClearStraight(qx, qy, kx, ky, board)
    }
    if abs(qx-kx) == abs(qy-ky) {
      return isPathClearDiagonal(qx, qy, kx, ky, board)
    }
    return false
}

func isPathClearStraight(qx, qy, kx, ky int, board [][]int) bool {
  if qx == kx {
    start, end := min(qy, ky)+1, max(qy, ky)
    for y := start; y < end; y++ {
      if board[qx][y] != 0 {
        return false
      }
    }
    return true
  }
  if qy  == ky {
    start, end := min(qx, kx)+1, max(qx,kx)
    for x := start; x< end; x++{
      if board[x][qy] != 0 {
        return false
      }
    }
    return true
  }
  return false
}

func isPathClearDiagonal(qx, qy, kx, ky int, board [][]int) bool {
  xDir := 1
  if kx < qx {
    xDir = -1
  }
  yDir := 1
  if ky < qy {
    yDir = -1
  }

  x, y := qx+xDir, qy+yDir
  for x != kx && y != ky {
    if board[x][y] != 0 {
      return false
    }
    x += xDir
    y += yDir
  }
  return true
}
func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func min(a, b int) int {
  if a < b {
    return a
  }
  return b
}

func max(a,b int) int {
  if a > b {
    return a
  }
  return b
}

func convertChess(s string) (int, int) {
  a, b := rune(s[0]), rune(s[1])
  x := int(a) - 97
  y := 7 - (int(b) - 49)
  return y, x
}
