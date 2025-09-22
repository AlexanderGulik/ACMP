package main

import (
	"fmt"
	"sort"
)

func main() {
	var a [5]int
	for i := 0; i < 5; i++ {
		fmt.Scan(&a[i])
	}
	sort.Ints(a[:])
	freq := make(map[int]int)
	for _, num := range a {
		freq[num]++
	}
	freqCount := make(map[int]int)

	for _, count := range freq {
		freqCount[count]++
	}

	switch {
	case freqCount[5] == 1:
		fmt.Println("Impossible")
	case freqCount[4] == 1:
		fmt.Println("Four of a Kind")
	case freqCount[3] == 1 && freqCount[2] == 1:
		fmt.Println("Full House")
	case isStaight(a):
		fmt.Println("Straight")
	case freqCount[3] == 1:
		fmt.Println("Three of a Kind")
	case freqCount[2] == 2:
		fmt.Println("Two Pairs")
	case freqCount[2] == 1:
		fmt.Println("One Pair")
	default:
		fmt.Println("Nothing")
	}

}

func isStaight(a [5]int) bool {
	isNormalStaright := true
	for i := 1; i < 5; i++ {
		if a[i] != a[i-1]+1 {
			isNormalStaright = false
			break
		}
	}
	isAceLowStr := a[0] == 1 && a[1] == 2 && a[2] == 3 && a[3] == 4 && a[4] == 5
	return isAceLowStr || isNormalStaright
}
