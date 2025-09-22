package main

import (
	"fmt"
	"math"
)

type Point struct {
	x, y float64
}

func main() {
	var n int
	fmt.Scan(&n)
	points := make([]Point, n)

	for i := 0; i < n; i++ {
		fmt.Scan(&points[i].x, &points[i].y)
	}
	var k int
	fmt.Scan(&k)
	if k == 0 {
		lenght := 0.0
		for i := 0; i < n; i++ {
			next := (i + 1) % n
			dx := points[next].x - points[i].x
			dy := points[next].y - points[i].y
			lenght += math.Sqrt(dx*dx + dy*dy)
		}
		fmt.Printf("%.10f", lenght)
		return
	}
	for iteration := 0; iteration < k; iteration++ {
		newPoints := make([]Point, n)
		for i := 0; i < n; i++ {
			next := (i + 1) % n
			newPoints[i].x = (points[i].x + points[next].x) / 2
			newPoints[i].y = (points[i].y + points[next].y) / 2
		}
		points = newPoints
	}
	lenght := 0.0
	for i := 0; i < n; i++ {
		next := (i + 1) % n
		dx := points[next].x - points[i].x
		dy := points[next].y - points[i].y
		lenght += math.Sqrt(dx*dx + dy*dy)
	}
	fmt.Printf("%.10f", lenght)
}
