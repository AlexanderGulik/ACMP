package main

import (
	"fmt"
)

func checkHubSize(points []int, hubSize int) bool {
	n := len(points) / hubSize
	for i := 0; i < n; i++ {
		start := i * hubSize
		for j := 0; j < hubSize; j++ {
			if points[j] != points[start+j] {
				return false
			}
		}
	}
	return true
}

func main() {
	var n int
	fmt.Scan(&n)
	n = n - 1
	points := make([]int, n)
	for i := 0; i < n; i++ {
		fmt.Scan(&points[i])
	}

	minHubSize := n
	for hubSize := 1; hubSize <= n; hubSize++ {
		if n%hubSize != 0 {
			continue
		}
		if checkHubSize(points, hubSize) {
			minHubSize = hubSize
			break
		}
	}

	fmt.Println(minHubSize)
}
