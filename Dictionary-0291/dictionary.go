package main
import (
  "fmt"
  "bufio"
  "os"
)

func main (){
  inputFile, _ := os.Open("input.txt")
  defer inputFile.Close()
  scanner := bufio.NewScanner(inputFile)
  var n int
  scanner.Scan()
  fmt.Sscan(scanner.Text(), &n)
  words := make([]string, n)
  for i := 0; i < n; i++ {
    scanner.Scan()
    words[i] = scanner.Text()
  }
  scanner.Scan()
  letters := scanner.Text()
  letterCount := countLetters(letters)
  count := 0
  for _, word := range words {
    wordCount := countLetters(word)
    canForm := true
    for char, cnt := range wordCount {
      if letterCount[char] < cnt {
        canForm = false
        break
      }
    }
    if canForm {
      count++
    }
  }
  fmt.Println(count)
}

func countLetters(s string) map[rune]int{
  countMap := make(map[rune]int)
  for _, char := range s {
    countMap[char]++
  }
  return countMap
}
