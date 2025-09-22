package main

import (
	"bufio"
	"fmt"
	"os"
)

type Point struct {
	x, y int
}

func main() {
	in, _ := os.Open("input.txt")
	out, _ := os.Create("output.txt")
	defer in.Close()
	defer out.Close()

	reader := bufio.NewReader(in)
	writer := bufio.NewWriter(out)
	defer writer.Flush()
	var n, m int
	fmt.Fscanf(reader, "%d %d\n", &n, &m)
	maze := make([]string, n)
	for i := 0; i < n; i++ {
		line, _ := reader.ReadString('\n')
		maze[i] = line[:len(line)-1]
	}

	if maze[n-2][m-2] == '@' {
		fmt.Fprintln(writer, -1)
		return
	}
	visits := make([][]int, n)
	for i := range visits {
		visits[i] = make([]int, m)
	}

	directions := []Point{{1, 0}, {0, 1}, {-1, 0}, {0, -1}}
	x, y := 1, 1
	visits[x][y] = 1
	steps := 0
	dir := 0
	for x != n-2 || y != m-2 {
		if steps > 10000000 {
			fmt.Fprintln(writer, -1)
			return
		}
		minVisits := -1
		candidateDirs := []int{}
		for d := 0; d < 4; d++ {
			nx, ny := x+directions[d].x, y+directions[d].y
			if maze[nx][ny] != '@' {
				if minVisits == -1 || visits[nx][ny] < minVisits {
					minVisits = visits[nx][ny]
					candidateDirs = []int{d}
				} else if visits[nx][ny] == minVisits {
					candidateDirs = append(candidateDirs, d)
				}
			}
		}
		newDir := -1

		for _, d := range candidateDirs {
			if d == dir {
				newDir = d
				break
			}
		}
		if newDir == -1 {
			for _, priorityDir := range []int{0, 1, 2, 3} {
				for _, d := range candidateDirs {
					if d == priorityDir {
						newDir = d
						break
					}
				}
				if newDir != -1 {
					break
				}
			}
		}
		dir = newDir
		x += directions[dir].x
		y += directions[dir].y
		visits[x][y]++
		steps++
	}
	fmt.Fprintln(writer, steps)
}
