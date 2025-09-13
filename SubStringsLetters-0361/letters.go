package main
import (
  "fmt"
)

func main(){
  var s string
  fmt.Scan(&s)
  result := 0

  for i := 0; i < len(s); i++{
    currLenght := 0
    currLetters := s[i]
    for j := i; j < len(s); j++{
     if currLetters == s[j] {
        if currLenght > result {
          result = currLenght
        }
     }
      currLenght++
    } 
  }
  fmt.Println(result)
}
