package main

import "fmt"

func main() {
	var n int32
	fmt.Scan(&n)
	
	if n <= 1 {
		fmt.Println(0)
	} else if n == 2 {
		fmt.Println(1)
	} else if n == 3 {
		fmt.Println(2)
	} else {
		fmt.Println(n)
	}
}
