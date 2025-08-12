package main

import (
	"fmt"
	"strings"
)

func solveFast(m, n int) string {
	var result strings.Builder
	
	for m >= 1 && n >= 2 {
		result.WriteString("BGG")
		m--
		n -= 2
	}
	
	for m > 0 {
		result.WriteString("B")
		m--
	}
	
	for n > 0 {
		result.WriteString("G")
		n--
	}
	
	return result.String()
}

func countActive(s string) int {
	n := len(s)
	result := 0
	for i := 0; i < n; i++ {
		prev := (i + n - 1) % n
		next := (i + 1) % n
		if s[i] == 'B' && s[prev] == 'G' && s[next] == 'G' {
			result++
		} else if s[i] == 'G' && (
			(s[prev] == 'B' && s[next] == 'G') || 
			(s[prev] == 'G' && s[next] == 'B')) {
			result++
		}
	}
	return result
}

func generateAllPermutations(b, g int) []string {
	s := strings.Repeat("B", b) + strings.Repeat("G", g)
	permutations := make(map[string]bool)
	permutations[s] = true
	
	runes := []rune(s)
	n := len(runes)
	for {
		var i, j int
		for i = n - 2; i >= 0; i-- {
			if runes[i] < runes[i+1] {
				break
			}
		}
		if i < 0 {
			break
		}
		for j = n - 1; j > i; j-- {
			if runes[j] > runes[i] {
				break
			}
		}
		runes[i], runes[j] = runes[j], runes[i]
		for k, l := i+1, n-1; k < l; k, l = k+1, l-1 {
			runes[k], runes[l] = runes[l], runes[k]
		}
		permutations[string(runes)] = true
	}
	
	var result []string
	for p := range permutations {
		result = append(result, p)
	}
	return result
}

func solveSlow(b, g int) int {
	if b+g < 3 {
		return 0
	}
	
	permutations := generateAllPermutations(b, g)
	maxActive := 0
	for _, p := range permutations {
		active := countActive(p)
		if active > maxActive {
			maxActive = active
		}
	}
	return maxActive
}

func compare() {
	for b := 0; b <= 10; b++ {
		for g := 0; g <= 10; g++ {
			if b+g < 3 {
				continue
			}
			slow := solveSlow(b, g)
			fast := countActive(solveFast(b, g))
			if slow != fast {
				fmt.Printf("Mismatch for b=%d, g=%d: slow=%d, fast=%d\n", b, g, slow, fast)
			}
		}
	}
}

func main() {
	var m, n int
	fmt.Scan(&m, &n)
	
	if m <= 10 && n <= 10 {
		compare()
	}
	
	result := solveFast(m, n)
	fmt.Println(result)
}

