package main

import (
  "fmt"
  "strconv"
  "strings"
)

func main(){
  var str1, str2 string
  fmt.Scan(&str1, &str2)
  result := countDay(str1, str2)
  fmt.Println(result)
}

func countDay(s1, s2 string) int {
  str1 := strings.Split(s1, ".")
  str2 := strings.Split(s2, ".")
  d1, _ := strconv.Atoi(str1[0])
  m1, _ := strconv.Atoi(str1[1])
  y1, _ := strconv.Atoi(str1[2])

  d2, _ := strconv.Atoi(str2[0])
  m2, _ := strconv.Atoi(str2[1])
  y2, _ := strconv.Atoi(str2[2])
  days1 := daysSince(d1, m1, y1)
  days2 := daysSince(d2, m2, y2)
  return days2 - days1 +1
}
func daysSince(d, m, y int) int {
  result := 0
  for i := 0; i < y; i++{
    if isVisokost(i){
      result += 366
    } else {
      result += 365
    }
  }
  for j := 1; j < m; j++{
    result += daysMonth(j, y)
  }
  result += d-1
  return result
}

func isVisokost(y int) bool {
  if y % 400 == 0 {
    return true
  }
  if y % 100 == 0 {
    return false
  }
  return y % 4 == 0
}

func daysMonth(m, y int) int {
  switch m {
    case 1, 3, 5, 7, 8, 10, 12:
      return 31
    case 4, 6, 9, 11:
      return 30
    case 2:
      if isVisokost(y){
        return 29
      }
      return 28
    default:
      return 0
  }
}
