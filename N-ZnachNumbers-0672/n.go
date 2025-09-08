package main

import (
	"fmt"
)
var tenPows []int64
func C(n, m int) int {
	if n-m > m {
		m = n - m
	}
	c := 1
	for i := n; i > m; i-- {
		c *= i
	}
	for i := n - m; i > 1; i-- {
		c /= i
	}
	return c
}

func cond(num int64) bool {
	sum := 0
	prod := 1
	for num > 0 {
		a := int(num % 10)
		sum += a
		prod *= a
		num /= 10
	}
	return sum == prod
}

func analyzeN(N int) {
	tenPows = make([]int64, N)
	tenPows[0] = 1
	for d := 1; d < N; d++ {
		tenPows[d] = 10 * tenPows[d-1]
	}

	count := 0
	var first int64 = 0
	num := int64(0)
	for d := 0; d < N; d++ {
		num += tenPows[d]
	}

	for {
		if cond(num) {
			if first == 0 {
				first = num
			}
			c := 1
			l := 1
			d := 1
			for ; d < N; d++ {
				if (num/tenPows[d])%10 == (num/tenPows[d-1])%10 {
					l++
				} else {
					c *= C(d, l)
					l = 1
				}
			}
			c *= C(d, l)
			count += c
		}
		d := 0
		for num%10 == 9 {
			num /= 10
			d++
		}
		if num == 0 {
			break
		}
		num++
		digit := num % 10
		for ; d > 0; d-- {
			num = num*10 + digit
		}
	}

	fmt.Printf("%d %d\n", count, first)
}

func main() {
	var N int
	fmt.Scan(&N)
	if N == 1 {
		fmt.Println("10 0")
		return
	}
	analyzeN(N)
}
/* мой код слишком медленный
package main

import (
    "fmt"
)

func main() {
    var n int
    fmt.Scan(&n)
    if n == 1 {
        fmt.Println("10 0")
        return
    }
    start := 1
    for i := 0; i < n-1; i++ {
        start *= 10
    }
    end := start*10 - 1
    count := 0
    minNum := -1
    
    for num := start; num <= end; num++ {
        if checkNumber(num) {
            count++
            if minNum == -1 || num < minNum {
                minNum = num
            }
        }
    }
    fmt.Printf("%d %d", count, minNum)
}

func checkNumber(num int) bool {
    sum, product := 0, 1
    temp := num
    for temp > 0 {
        digit := temp % 10
        sum += digit
        product *= digit
        temp /= 10
    }
    return sum == product
}
*/
