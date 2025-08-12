package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)
	countEmpty := 0
	j := 1
	k := n
	for n%4 != 0 {
		n++
		countEmpty++
	}
	s := 1
	for i := 0; k >= j; i++ {
		fmt.Printf("%d %d ", i/2+1, s)
		if s == 1 {
			if countEmpty > 0 {
				fmt.Printf("%d %d \n", 0, j)
        countEmpty--
			} else {
				fmt.Printf("%d %d \n", k, j)
				k--
			}
      j++
		} else {
			if countEmpty > 0 {
   				fmt.Printf("%d %d ", j, 0)
				  countEmpty--
			} else {
				fmt.Printf("%d %d ", j, k)
				k--
	  		}
        j++
			}
      if s == 1 {
				s = 2
			} else {
				s = 1
			}
	}
}



/*
int main(){
  int n;
  cin >>n;
  int count_empty = 0, j =1, k = n;
  while (n %4 !=0){
    n++;
    count_empty++;
  }
  int s = 1;
  for (int i = 0; k >= j; i++){
    count << i /2+1 << ' ' << s << ' ';
    if (s == 1)
    cout << (count_empty-- > 0  0 : k--) << ' '<< j++ << endl;
    else
    cout << j++ << ' ' << (count_empty-- > 0 ? 0 : k--) << endl;
    s = (s == 1 ? 2 : 1)
  }
  return 0
}
*/
