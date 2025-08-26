package main

import (
	"fmt"
)

type BitSet []uint64

func (b BitSet) Set(pos int) {
	b[pos>>6] |= 1 << (pos & 63)
}

func (b BitSet) Get(pos int) bool {
	return (b[pos>>6] & (1 << (pos & 63))) != 0
}

func main() {
	var n, m int
	fmt.Scan(&n, &m)
	var k int
	fmt.Scan(&k)
	totalCell := n * m

	bits := (totalCell + 63) >> 6 
	model := make(BitSet, bits)
	infected := 0

	queue := make([]int, 0, k)
	for i := 0; i < k; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		x--
		y--
		pos := x*m + y
		if !model.Get(pos) {
			model.Set(pos)
			infected++
			queue = append(queue, pos)
		}
	}

	days := 0
	nextQueue := make([]int, 0, k*4)
	for infected < totalCell && len(queue) > 0 {
		days++
		for _, cell := range queue {
			x, y := cell/m, cell%m
			if x > 0 {
				tryInfect(model, &infected, cell-m, &nextQueue, totalCell)
			}
			if x < n-1 {
				tryInfect(model, &infected, cell+m, &nextQueue, totalCell)
			}
			if y > 0 {
				tryInfect(model, &infected, cell-1, &nextQueue, totalCell)
			}
			if y < m-1 {
				tryInfect(model, &infected, cell+1, &nextQueue, totalCell)
			}
		}
		queue, nextQueue = nextQueue, queue[:0]
	}

	fmt.Println(days)
}

func tryInfect(model BitSet, infected *int, pos int, queue *[]int, totalCell int) {
	if !model.Get(pos) {
		model.Set(pos)
		*infected++
		*queue = append(*queue, pos)
	}
}
/*просто ебал реп огранечений
package main
import (
  "fmt"
)

func main(){
  var n, m int
  fmt.Scan(&n, &m)
  var k int
  fmt.Scan(&k)
  lenK := k*2
  model := make([]byte, n*m)
  infected := 0
  totalCell := n*m
  queue := make([]int, 0, k)
  for i := 0; i < lenK; i += 2{
    var x, y int
    fmt.Scan(&x, &y)
    x--
    y--
    pos := x*m +y
    if model[pos] == 0{
       model[pos] = 1
       infected++
       queue = append(queue, pos)
    }
  } 
  days := 0
  for infected < totalCell && days < 3000 {
    days++
    nextQueue := make([]int, 0, len(queue)*4)
    for _, cell := range queue {
      x, y := cell/m, cell%m
        if x > 0 { tryInfect(model, &infected, cell-m, &nextQueue, totalCell) }
        if x < n-1 { tryInfect(model, &infected, cell+m, &nextQueue, totalCell) }
        if y > 0 { tryInfect(model, &infected, cell-1, &nextQueue, totalCell) }
        if y < m-1 { tryInfect(model, &infected, cell+1, &nextQueue, totalCell) }
        if infected == totalCell { break }
      }
       queue = nextQueue
      if infected == totalCell { break }
    } 

  fmt.Println(days)
}

func tryInfect( model []byte, infected *int, pos int, queue *[]int, totalCell int) {
  if model[pos] == 0 {
    model[pos] = 1
    *infected++
    *queue = append(*queue, pos)
  }
}
*/


