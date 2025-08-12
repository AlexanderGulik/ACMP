package main

import (
	"bufio"
	"os"
	"strings"
)

func main() {

	inputFile, _ := os.Open("INPUT.TXT")
	defer inputFile.Close()
	scanner := bufio.NewScanner(inputFile)
	scanner.Scan()
	encoded := scanner.Text()
  var decoded strings.Builder
  a := make(map[rune]int)
  k := 0
  for i := '0'; i <= '9'; i++{
    a[i] = k
    k++
  }
  for i := 'A'; i <= 'Q'; i++{
    a[i] = k
    k++
  }
  b := make([]rune, 0, 100)
  b = append(b, ' ')
  for i := 'a'; i <= 'z'; i++{
    b = append(b, i)
  }
  for i, c := range encoded{
    p := a[c]
    w := ((p - i - 1) % 27 +27) %27
    decoded.WriteRune(b[w])
  }
	outputFile, _ := os.Create("OUTPUT.TXT")
	defer outputFile.Close()
	outputFile.WriteString(decoded.String())
}


