package main

import (
	"fmt"
)

type Ingredient struct {
	sign    string
	amount  int
	precent float64
}

func main() {
	var n int
	fmt.Scan(&n)
	product := make([]Ingredient, n)
	sum := 0
	for i := 0; i < n; i++ {
		var a string
		var b int
		fmt.Scan(&a)
		fmt.Scan(&b)
		product[i] = Ingredient{a, b, 0.0}
		sum += b
	}

	for i := 0; i < n; i++ {
		product[i].precent = float64(product[i].amount) * 100 / float64(sum)
	}

	result := make([]int, n)
	total := 0
	remainders := make([]float64, n)
	for i := 0; i < n; i++ {
		result[i] = int(product[i].precent)
		remainders[i] = product[i].precent - float64(result[i])
		total += result[i]
	}
	remaing := 100 - total

	for remaing > 0 {
		bestIndx := -1
		bestRemainder := -1.0
		for i := 0; i < n; i++ {
			if product[i].sign == "+" {
				if remainders[i] > bestRemainder {
					bestRemainder = remainders[i]
					bestIndx = i
				}
			}
		}
		if bestIndx == -1 {
			for i := 0; i < n; i++ {
				if product[i].sign == "-" && remainders[i] > bestRemainder {
					bestRemainder = remainders[i]
					bestIndx = i
				}
			}
		}
		if bestIndx != -1 {
			result[bestIndx]++
			remainders[bestIndx] = -1
			remaing--
		} else {
			break
		}
	}
	for i := 0; i < n; i++ {
		fmt.Println(result[i])
	}
}

/*#include <fstream>

#include <vector>



using namespace std;



int main() {

    ifstream in("input.txt");

    ofstream out("output.txt");

    int n;      // количество ингредиентов

    in >> n;

    int plus = 0;           // из них - положительных

    int minus = 0;          // отрицательных

    int neyttral = 0;       // нельзя менять

    int massa = 0;        // общая масса

    int prOk = 0;         // сумма процентов округленных в меньшую сторону



    vector <int> ingr(n); // Собственно, доли ингредиентов

    vector <int> okrugl(n);   // округленные в меньшую сторону проценты

    vector <char> p(n);   // признак полезности





    for (int z = 0; z < n; z++) { // грузим все и считаем общую массу

        in >> p[z];

        in >> ingr[z];

        massa += ingr[z];

    }

    for (int z = 0; z < n; z++) { // считаем проценты

        okrugl[z] = (ingr[z] * 100) / massa;   // округленные в меньшую сторону проценты

        prOk += okrugl[z];

        if ((ingr[z] * 100) % massa == 0) {

            p[z] = '0';        // этот ингредиент менять нельзя

            continue;

        }

        if (p[z] == '+') plus++;  // считаем полезные и вредные

        else minus++;

    }

    for (int z = 0; z < n; z++) { // выводим результат

        if (p[z] == '-' && prOk < 100 && plus < minus) {

            out << okrugl[z] + 1 << endl; // даже если все вредное, сумма процентов должна быть не менее 100

            prOk++;     // улучшаем баланс

            minus--;    // пока не уравняем число вредных и полезных

        }

        else if (p[z] == '0' || p[z] == '-') out << okrugl[z] << endl;  // здесь не увеличиваем однозначно

            else if (p[z] == '+' && prOk < 100) {  // а здесь увеличиваем, ибо полезно

                out << okrugl[z] + 1 << endl;

                prOk++;  // и улучшаем баланс

            }

            else out << okrugl[z] << endl; // мы вышли на 100%, дальше выдаем по нижней границе

    }

    return 0;

}

*/
