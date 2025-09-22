package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)
	dp := make([][]int, n+1)
	for i := range dp {
		dp[i] = make([]int, k)
	}
	for digit := 0; digit < k; digit++ {
		if digit != 0 {
			dp[1][digit] = 1
		}
	}
	for lenght := 2; lenght <= n; lenght++ {
		for lastDigit := 0; lastDigit < k; lastDigit++ {
			for prevDigit := 0; prevDigit < k; prevDigit++ {
				if prevDigit != 0 || lastDigit != 0 {
					dp[lenght][lastDigit] += dp[lenght-1][prevDigit]
				}
			}
		}
	}
	total := 0
	for digit := 0; digit < k; digit++ {
		total += dp[n][digit]
	}
	fmt.Println(total)
}
