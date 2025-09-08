


//ебал рот разрабов
/*#include <iostream>
#include <string>
using namespace std;
int jul(int year, int month, int day) {
       // вычисление номера дня юлианского календаря по григорианской дате
       int a = (14 - month) / 12;
       int y = year + 4800 - a;
       int m = month + 12 * a - 3;
       return day + (153 * m + 2) / 5 + 365 * y + y / 4 - y / 100 + y / 400 - 32045;
}

int main() {
       int dr, mr, dt, mt, yt;
       cin >> dr >> mr >> dt >> mt >> yt;
       if (dr == dt && mr == mt) { cout << 0; return 0; }
       if (!(dr == 29 && mr == 2)) {
             if (jul(yt, mt, dt) > jul(yt, mr, dr)) // ДР уже прошел
                    cout << jul(yt + 1, mr, dr) - jul(yt, mt, dt);
             else cout << jul(yt, mr, dr) - jul(yt, mt, dt);
       }
       else {
             int yr = yt;
             while (!(yr % 400 == 0 || (yr % 4 == 0 && yr % 100 != 0))) yr++;
             if (jul(yt, mt, dt) > jul(yr, mr, dr)) {   // ДР уже прошел
                    yr++;
                    while (!(yr % 400 == 0 || (yr % 4 == 0 && yr % 100 != 0))) yr++;
                    cout << jul(yr, mr, dr) - jul(yt, mt, dt);
             }
             else cout << jul(yr, mr, dr) - jul(yt, mt, dt);
       }
       return 0;
}*/
/*package main

import (
	"fmt"
	"time"
)

func isLeap(y int) bool {
	return y%400 == 0 || (y%4 == 0 && y%100 != 0)
}

func main() {
	var dr, mr int
	fmt.Scan(&dr, &mr)
	var dt, mt, yt int
	fmt.Scan(&dt, &mt, &yt)

	today := time.Date(yt, time.Month(mt), dt, 0, 0, 0, 0, time.UTC)

	if dr == dt && mr == mt {
		fmt.Println(0)
		return
	}


	if dr == 29 && mr == 2 {
		yr := yt
		for !isLeap(yr) {
			yr++
		}
		birthday := time.Date(yr, time.February, 29, 0, 0, 0, 0, time.UTC)
		if !birthday.After(today) { 
			yr++
			for !isLeap(yr) {
				yr++
			}
			birthday = time.Date(yr, time.February, 29, 0, 0, 0, 0, time.UTC)
		}
		days := int(birthday.Sub(today) / (24 * time.Hour))
		fmt.Println(days)
		return
	}


	birthday := time.Date(yt, time.Month(mr), dr, 0, 0, 0, 0, time.UTC)
	if !birthday.After(today) { 
		birthday = time.Date(yt+1, time.Month(mr), dr, 0, 0, 0, 0, time.UTC)
	}
	days := int(birthday.Sub(today) / (24 * time.Hour))
	fmt.Println(days)
}*/
/*package main
import (
  "fmt"
)

func jul(y, m, d int) int {
  a := (14-m)/12
  year := y +4800 -a
  month := m +12*a-3
  return d + (153*month+2)/5 + 365*year + year/4 - year/100 + year/400 - 32045
}

func isLeap(y int) bool {
	if y%400 == 0 {
		return true
	}
	if y%100 == 0 {
		return false
	}
	return y%4 == 0
}

func main(){
  var dr, mr int
  fmt.Scan(&dr, &mr)
  var dt, mt, yt int
  fmt.Scan(&dt, &mt, &yt)
  if dr == dt && mr == mt {
    fmt.Println(0)
    return
  }
  if !(dr ==29 && mr ==2){
    if jul(yt, mt, dt) > jul(yt, mr, dr) {
      fmt.Println(jul(yt+1, mr, dr) - jul(yt,mt,dt))
    } else {
      fmt.Println(jul(yt, mr, dr) - jul(yt,mt,dt))
    }
  } else {
    yr := yt
    for !isLeap(yr){
      yr++
    }
    if jul(yt, mt, dt) > jul(yr,mr,dr){
      yr++
      for !isLeap(yr){
        yr++
      }
      fmt.Println(jul(yr, mr, dr) - jul(yt, mt, dt))
    } else {
      fmt.Println(jul(yr, mr, dr) - jul(yt, mt, dt))
    }
  }
}


*/



/*package main

import (
  "fmt"
)

func main(){
  var dayB, monthB int
  fmt.Scan(&dayB, &monthB)
  var currDay, currMonth, curryear int
  fmt.Scan(&currDay, &currMonth, &curryear)
  if dayB == currDay && currMonth == monthB {
    fmt.Println(0)
    return
  }
  targetYear := curryear

  if monthB < currMonth || (monthB == currMonth && dayB < currDay){
    targetYear = curryear+1
  }

  if dayB == 29 && monthB == 2{
    for !isVisokost(targetYear){
      targetYear++
    }
  }
  result := 0
  if targetYear == curryear {
     result = dayOfYear(dayB, monthB, curryear) - dayOfYear(currDay, currMonth, curryear)    
  } else {
    result = daysInYear(curryear) - dayOfYear(currDay, currMonth, curryear)
    for y := curryear+1 ; y < targetYear; y++{
      result += daysInYear(y)
    }
    result += dayOfYear(dayB, monthB,targetYear)
  }
  fmt.Println(result)
}

func dayOfYear(d, m, y int) int{
  days := 0
  for i := 1; i < m; i++{
    days += daysMonth(i, y)
  }
  return days+d
}
func daysInYear(y int) int {
  if isVisokost(y){
    return 366
  }
  return 365
}
func isVisokost(y int) bool {
  if y % 400 == 0 {
    return true
  }
  if y % 100 == 0 {
    return false
  }
  return y % 4 == 0
}

func daysMonth(m, y int) int {
  switch m {
    case 1, 3, 5, 7, 8, 10, 12:
      return 31
    case 4, 6, 9, 11:
      return 30
    case 2:
      if isVisokost(y){
      return 29
      }
      return 28
    default:
      return 0
  }
}*/
