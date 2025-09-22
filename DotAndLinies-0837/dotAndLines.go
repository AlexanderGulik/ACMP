package main

import "fmt"

func main() {
	var k int
	fmt.Scan(&k)
	for i := 0; i < k; i++ {
		var n, m int
		fmt.Scan(&n, &m)
		if n == 1 {
			fmt.Println("No")
			continue
		}
		maxPos := n * (n - 1) / 2
		if m > maxPos {
			fmt.Println("No")
			continue
		}

		maxEdge := (n - 1) * (n - 2) / 2
		if m <= maxEdge {
			fmt.Println("Yes")
		} else {
			fmt.Println("No")
		}

	}
}
