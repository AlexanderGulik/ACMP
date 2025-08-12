package main
import (
  "os"
  "fmt"
  "strings"
  "strconv"
  "bufio"
)
func main () {
  content, _ := os.ReadFile("input.txt")
  nums := strings.Split(strings.TrimSpace(string(content)), " ")
  result := 0
  start, _ := strconv.Atoi(nums[0]) 
  end, _ :=  strconv.Atoi(nums[1])
  for i:= start; i <= end; i++{
    num :=  prdocut(i)
    fmt.Println(i, " non ", num)
    if num != 0 && i % num == 0 {
    fmt.Println(i, " Yes ", num)
      result++
    }
  }
  outFile, _ := os.Create("output.txt")
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)
  fmt.Fprintln(writer, result)
  writer.Flush()
  
}
func prdocut(num int) int {
  res := 1
  strNum := strconv.Itoa(num)
  for i := 0; i < len(strNum); i++{
    digit, _ := strconv.Atoi(string(strNum[i]))
    res *= digit
  }
  return res
}


