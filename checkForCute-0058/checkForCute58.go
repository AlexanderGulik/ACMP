package main 

import (

	"fmt"
	"os"
  "strconv"
	"strings"
  "bufio"
)

func main () {
  inFile, _ := os.Open("input.txt")
  defer inFile.Close()
  scanner := bufio.NewScanner(inFile)
  scanner.Scan()
  t, _ := strconv.Atoi(scanner.Text())
  outputs := make([]string, 0, t)
  for i := 0; i < t; i++ {
    scanner.Scan()
    sizes := strings.Fields(scanner.Text())
    n, _ := strconv.Atoi(sizes[0])
    m, _ := strconv.Atoi(sizes[1])
    matrix := make([][]int, n)
    for j := 0; j < n; j++ {
      scanner.Scan()
      nums := strings.Fields(scanner.Text())
      row := make([]int, m)
      for k := 0; k < m; k++{
        row[k], _ = strconv.Atoi(nums[k])
      }
      matrix[j] = row
    }

    if isCute(matrix, n, m) {
      outputs = append(outputs, "YES")
    } else {
      outputs = append(outputs, "NO")
    }
  }
  outFile, _ := os.Create("OUTPUT.TXT")
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)
  for _, line := range outputs {
    fmt.Fprintln(writer, line)
  }
  writer.Flush()
}

func isCute(matrix [][]int, n, m int) bool {
  if n < 2 || m < 2{
    return true
  }
  for i := 0; i < n-1; i++{
    for j := 0; j < m-1; j++{
      a := matrix[i][j]
      b := matrix[i+1][j]
      c := matrix[i][j+1]
      d := matrix[i+1][j+1]
      if a == b && a == c && a == d{
        return false
      }
    }
  }
  return true
}
