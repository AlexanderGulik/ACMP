package main
import (
    "fmt"
  )

func main(){
 var n int
 fmt.Scan(&n)
 a := []int{2,3,4,7,13,15}
 b := []int{1,5,6,8,9,10,11,12,14}
 if n > len(a){
   for i := len(a); i <= n; i++{
   nextA :=  b[i-1] + b[i-3]
   a = append(a, nextA)
   lastB := b[len(b)-1]
     for candidate := lastB+2; candidate < nextA; candidate++{
       b =append(b, candidate)
     }
   }
 }

 fmt.Println(a[n-1])
 fmt.Println(b[n-1])
}

