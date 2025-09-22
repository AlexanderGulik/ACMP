package main

import (
	"fmt"
	"math"
)

func main() {
	var n int
	fmt.Scan(&n)
	tables := make([][]int, n)
	for i := 0; i < n; i++ {
		tables[i] = make([]int, n)
	}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Scan(&tables[i][j])
		}
	}
	result := math.MinInt32
	dirs := [][2]int{{-1, 0}, {1, 0}, {0, -1}, {0, 1}}
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			b := tables[i][j]
			for _, d1 := range dirs {
				ai, aj := i+d1[0], j+d1[1]
				if ai < 0 || ai >= n || aj < 0 || aj >= n {
					continue
				}
				a := tables[ai][aj]
				for _, d2 := range dirs {
					bi, bj := i+d2[0], j+d2[1]
					if bi < 0 || bi >= n || bj < 0 || bj >= n || (bi == ai && bj == aj) {
						continue
					}
					c := tables[bi][bj]
					result = max(result, a+b+c)
				}
			}
		}
	}
	fmt.Println(result)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
