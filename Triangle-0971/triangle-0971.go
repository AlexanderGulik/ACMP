package main

import (
	"fmt"
)

func main() {
	var x, y int64
	fmt.Scan(&x, &y)

	ax, ay := x-1, y+1
	bx, by := x+1, y+1
	cx, cy := x, y-2

	coords := []int64{ax, ay, bx, by, cx, cy}
	for _, v := range coords {
		if  v < -1_000_000_000 || v > 1_000_000_000 {
			fmt.Println("NO")
			return
		}
	}

	vx1 := bx - ax
	vy1 := by - ay
	vx2 := cx - ax
	vy2 := cy - ay
	if vx1*vy2 - vy1*vx2 == 0 {
		fmt.Println("NO")
		return
	}
	fmt.Println("YES")
	fmt.Println(ax, ay)
	fmt.Println(bx, by)
	fmt.Println(cx, cy)
}





