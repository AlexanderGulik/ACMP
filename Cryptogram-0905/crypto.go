package main

import (
  "fmt"
  "os"
  "bufio"
  "strings"
  "strconv"
)

func main() {
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())
  descrypt := "the quick brown fox jumps over the lazy dog"
  var enscrypt string
  line := make([]string,n)
  for i := 0; i < n; i++{
    scanner.Scan()
    line[i] = scanner.Text()
    if checkString(line[i], descrypt){
      enscrypt = line[i]
    }
  }

  if enscrypt == ""{
    fmt.Println("No solution")  
    return
  }

    maps := make(map[byte]byte)
    reverseMap := make(map[byte]byte)

    str1 := strings.ReplaceAll(descrypt, " ", "")
    str2 := strings.ReplaceAll(enscrypt, " ", "")

    for i := 0; i< len(str1); i++{
      if left, right := maps[str2[i]]; right && left != str1[i]{
        fmt.Println("No solution")
        return
      } 
      if left, right := reverseMap[str1[i]]; right && left != str2[i]{
        fmt.Println("No solution")
        return
      }

      maps[str2[i]] = str1[i]
      reverseMap[str1[i]] = str2[i]
    }
    if len(maps) != 26{
      fmt.Println("No solution")  
     return
    }

    for i := 0; i < n; i++{
      var descrypted strings.Builder
      for j := 0; j< len(line[i]); j++{
        char := line[i][j]
        if char == ' '{
          descrypted.WriteByte(' ')
        } else {
          if dec, exit := maps[char]; exit {
            descrypted.WriteByte(dec)
          } else {
            descrypted.WriteByte(char)
          }
        }
      }
      fmt.Println(descrypted.String())
    }
}

func checkString(s, des string) bool {
  if len(s) != len(des){
    return false
  }
  arr1 := strings.Split(s, " ")
  arr2 := strings.Split(des, " ")
  if len(arr2) != len(arr1) {
    return false
  }
  for i := 0; i < len(arr1); i++{
    if len(arr1[i]) != len(arr2[i]){
      return false
    }
  }
  return true
}

