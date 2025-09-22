package main

import (
	"fmt"
	"strings"
)

func main() {
	var s string
	fmt.Scan(&s)
	reuslt := strings.Builder{}
	if strings.Contains(s, "_") {
		curr := strings.Split(s, "_")
		for i := 0; i < len(curr); i++ {
			if !isGood(curr[i]) {
				fmt.Println("Error!")
				return
			}
			if i == 0 {
				reuslt.WriteString(curr[i])
			} else {
				firstUpper := strings.ToUpper(curr[i][:1]) + curr[i][1:]
				reuslt.WriteString(firstUpper)
			}
		}
	} else {
		if isGood(string(s[0])) {
			reuslt.WriteString(string(s[0]))
		} else {
			fmt.Println("Error!")
			return
		}
		for i := 1; i < len(s); i++ {
			if s[i] >= 'A' && s[i] <= 'Z' {
				lower := string(s[i] + 32)
				reuslt.WriteString("_")
				reuslt.WriteString(lower)
			} else {
				reuslt.WriteString(string(s[i]))
			}
		}
	}
	fmt.Println(reuslt.String())
}

func isGood(s string) bool {
	if len(s) == 0 {
		return false
	}
	for i := 0; i < len(s); i++ {
		if !(s[i] >= 'a' && s[i] <= 'z') {
			return false
		}
	}
	return true
}
