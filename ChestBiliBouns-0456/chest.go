package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Scan(&x, &y)
	fib := make([]int, x+1)
	fib[1], fib[2] = 1, 1
	for i := 3; i <= x; i++ {
		fib[i] = fib[i-2] + fib[i-1]
	}

	b := 0
	for (y-fib[x-1]*b)%fib[x-2] != 0 {
		b++
	}

	a := (y - fib[x-1]*b) / fib[x-2]
	fmt.Println(a, b)

}

/*

#include <iostream>
#include <vector>
using namespace std;

int main() {
       int x, y;
       cin >> x >> y;
       vector <int> f(x+1); // Ряд Фиббоначчи
       f[1] = 1; f[2] = 1;
       for (int z = 3; z <= x; z++) f[z] = f[z - 2] + f[z - 1];
       int a, b;
       b = 0;
       while ((y - f[x - 1] * b) % f[x - 2] != 0) b++;       // подбираем сумму второго года
       a = (y - f[x - 1] * b) / f[x - 2];      // и вычисляем исодную сумму
       cout << a << " " << b;
       return 0;
}
*/
/*
func main() {
	var x, y int
	fmt.Scan(&x, &y)
	for a := 1; a <= y; a++ {
		for b := 1; b < y; b++ {
			if calucate(x, a, b) == y {
				fmt.Println(a, b)
				return
			}
		}
	}
}

func calucate(x, a, b int) int {
	if x == 1 {
		return a
	}

	if x == 2 {
		return b
	}

	prev2, prev1 := a, b
	for i := 3; i <= x; i++ {
		curr := prev2 + prev1
		prev2, prev1 = prev1, curr
	}
	return prev1
}*/

/* сначала он положил некоторое кол-во монет 
на второй год он вынул из сундука некоторое + кол-во монет
на третий год он добавил столько монет, сколько было 2 года назад = т. e. из первого

*/
