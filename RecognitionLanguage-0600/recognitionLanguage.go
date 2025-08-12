package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Scan()
	count, _ := strconv.Atoi(scanner.Text())
	outputs := make([]string, 0, count)
	for i := 0; i < count && scanner.Scan(); i++ {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			i--
			continue
		}
		if recognitionLanguage(line) {
			outputs = append(outputs, "YES")
		} else {
			outputs = append(outputs, "NO")
		}
	}
	outFile, err := os.Create("output.txt")
	if err != nil {
		panic(err)
	}
	defer outFile.Close()
	writer := bufio.NewWriter(outFile)
	for _, line := range outputs {
		fmt.Fprintln(writer, line)
	}
	writer.Flush()
}

func recognitionLanguage(s string) bool {
	if len(s) == 0 {
		return false
	}
	for _, c := range s {
		if c != '0' && c != '1' && c != '2' {
			return false
		}
	}
	n0, n1, n2 := 0, 0, 0
	phase := 0 
	for _, c := range s {
		switch phase {
		case 0:
			if c == '0' {
				n0++
			} else if c == '1' {
				phase = 1
				n1++
			} else {
				return false
			}
		case 1:
			if c == '1' {
				n1++
			} else if c == '2' {
				phase = 2
				n2++
			} else {
				return false
			}
		case 2:
			if c == '2' {
				n2++
			} else {
				return false
			}
		}
	}
	return n0 == n1 && n1 == n2 && n0 > 0
}
