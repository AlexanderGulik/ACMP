package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)
	result := make([]byte, n)
	for i := 0; i < n; i++ {
		var x, y, s int
		fmt.Scan(&x, &y, &s)
		if testFrame(x, y, s) {
			result[i] = '1'
		} else {
			result[i] = '0'
		}
	}
	fmt.Println(string(result))
}

func testFrame(x, y, s int) bool {
	if s == 1 || s == 2 {
		return true
	}

	if s > x || s > y {
		return false
	}

	if x%s == 0 {
		if (y-2)%s == 0 {
			return true
		}
		return false
	}

	if x%s == 1 {
		if (y-1)%s == 0 || ((y-2)%s == 0 && y%s == 0) {
			return true
		}
		return false
	}
	if x%s == 2 {
		if y%s == 0 {
			return true
		}
		return false
	}
	return false
}
