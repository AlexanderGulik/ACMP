package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	var k, n int
	fmt.Sscanf(scanner.Text(), "%d %d", &k, &n)
	
	s := make([]string, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		s[i] = scanner.Text()
		s[i] = strings.TrimSpace(s[i])
		
		if len(s[i]) > k {
			fmt.Println("Impossible.")
			return
		}
		
		w := k - len(s[i])
		left := w / 2
		right := w - left
		
		s[i] = strings.Repeat(" ", left) + s[i] + strings.Repeat(" ", right)
	}
	
	for i := 0; i < n; i++ {
		fmt.Println(s[i])
	}
}
