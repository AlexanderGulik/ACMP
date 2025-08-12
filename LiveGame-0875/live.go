package main
import (
  "fmt"
  "os"
  "bufio"
  "strings"
)

func main (){
  input, _ := os.Open("input.txt")
  defer input.Close()
  scanner := bufio.NewScanner(input)
  scanner.Scan()
  firstLine := scanner.Text()
  var n, m, k int
  fmt.Sscanf(firstLine, "%d %d %d", &n, &m, &k)
  game := make([][]rune, n)
  for i := 0; i< n && scanner.Scan(); i++{
    line := strings.TrimSpace(scanner.Text())
    game[i] = make([]rune, m)
    for j :=0; j<m; j++{
      if j < len(line) {
        game[i][j] = rune(line[j])
      } else {
        game[i][j] = '.'
      }
    }
  }
 for z:= 0; z< k; z++{
   result := make([][]rune, n)
   for i := range result {
    result[i] = make([]rune, m)
    copy(result[i], game[i])
   }
    for i := 0; i < n; i++{
      for j:= 0; j<m; j++{
        if game[i][j] == '.' {
          if checkForDie(game, i, j) {
            result[i][j] = '*'
          } else {
            result[i][j] = '.'
          }
      } else if game[i][j] == '*' {
        if checkForLive(game, i, j) {
          result[i][j] = '*'
        } else {
          result[i][j] = '.'
        }
      }
      }
    }
    game = result
 }
 file, _ := os.Create("OUTPUT.TXT")
 defer file.Close()
 for i := 0; i< n; i++{
    for j:= 0; j< m; j++{
      fmt.Fprintf(file, "%c", game[i][j])
    }
    fmt.Fprintln(file)
 }
}

func checkForLive(matrix [][]rune, i, j int) bool {
  count := 0
  cols := len(matrix[0])
  rows := len(matrix)
  direction := [8][2]int{
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1}, {0,1},
    {1, -1}, {1, 0}, {1, 1},
  }
  for _, dir := range direction {
    ni, nj := (i+dir[0]+rows)%rows, (j+dir[1]+cols)%cols
    if ni >=0 && ni <rows && nj >= 0 && nj <cols {
      if matrix[ni][nj] == '*' {
        count++
      }
    }
  }
  return count == 2|| count ==3
  }
func checkForDie(matrix [][]rune, i, j int) bool {
  count := 0
  cols := len(matrix[0])
  rows := len(matrix)
  direction := [8][2]int{
    {-1, -1}, {-1, 0}, {-1, 1},
    {0, -1}, {0,1},
    {1, -1}, {1, 0}, {1, 1},
  }
  for _, dir := range direction {
    ni, nj := (i+dir[0]+rows)%rows, (j+dir[1]+cols)%cols
    if ni >=0 && ni <rows && nj >= 0 && nj <cols {
      if matrix[ni][nj] == '*' {
        count++
      }
    }
  }
  return count == 3
}
