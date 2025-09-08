package main

import (
  "fmt"
  "bufio"
  "os"
)

func main(){
  sum := 0
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  scanner.Split(bufio.ScanRunes)
  for scanner.Scan() {
    char := rune(scanner.Text()[0])
    if (char >= 'a'&& char <= 'z') {
      num := int(char - 'a' +1)
      sum += DigitSum(num)
    } else if (char >= 'A'&& char <= 'Z') {
      num := int(char - 'A' +1)
      sum += DigitSum(num)+10
    } else if char == ' ' {
      sum +=4
    } else if (char >='0' && char <= '9'){
      digit := int(char -'0')
      sum += 13-digit
    } else if char == '.' {
      sum+=5
    } else if char == ';' {
      sum += 7
    } else if char == ',' {
      sum += 2
    } else if char == '=' || char == '+' || char == '-' || char =='\'' || char == '"' {
      sum += 3
    } else if char == '(' || char == ')' {
      sum++
    } else if char == '{' || char == '}' || char == '[' || char == ']' || char == '<' || char == '>' {
      sum +=8
    }

  }
  fmt.Println(sum)

}

func DigitSum(n int) int {
   sum := 0
   for n > 0 {
    sum += n%10
    n /=10
   }
   return sum
}
