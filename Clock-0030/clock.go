package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main(){
  var st1, st2 string
  fmt.Scan(&st1, &st2)
  start := timeDecomposition(st1)
  end := timeDecomposition(st2)
  counts := make([]int, 10)
  for sec := start; sec <= end; sec++{
    timeStr := secondsToTime(sec)
    countDigits(timeStr, counts)
  }
  for i := 0; i < 10; i++{
    fmt.Println(counts[i])
  }
}

func timeDecomposition(str string) int {
  parts := strings.Split(str, ":")
  h, _ := strconv.Atoi(parts[0])
  m, _ := strconv.Atoi(parts[1])
  s, _ := strconv.Atoi(parts[2])
  return (h*3600+m*60+s)
}

func secondsToTime(totalSec int) string{
  h := totalSec/ 3600
  remain := totalSec %3600
  m := remain / 60
  s := remain % 60
  return fmt.Sprintf("%02d:%02d:%02d", h, m, s)
}

func countDigits(timeStr string, counts[]int) {
    cleaned := strings.ReplaceAll(timeStr, ":", "")
    for _, char := range cleaned {
      digit := int(char - '0')
      if digit >= 0 && digit <=9 {
      counts[digit]++
    }
    }
}
