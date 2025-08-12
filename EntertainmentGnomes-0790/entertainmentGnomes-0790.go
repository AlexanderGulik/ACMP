package main 
import (
  "strconv"
  "strings"
  "os"
)

func main () {
  content, _ := os.ReadFile("input.txt")
  data := strings.Split(strings.TrimSpace(string(content)), "/")
  day, _ := strconv.Atoi(data[0])
  month, _ := strconv.Atoi(data[1])
  year, _ := strconv.Atoi(data[2])
  notation := day+1
  newDay :=  notationNumber(int(day), notation)
  newMonth := notationNumber(int(month), notation)
  newYear := notationNumber(int(year), notation)
  var result string
  result = newDay + "/" + newMonth + "/"+ newYear
  _ = os.WriteFile("output.txt", []byte(result), 0644)
}

func notationNumber(num, base int) string {
  cf := "0123456789ABCDEFGHIJKLMNOPQRSTUVWXYZ"
  var result string
  for num > 0 {
    result = string(cf[num%base]) + result
    num /= base
  }
  return result
}
