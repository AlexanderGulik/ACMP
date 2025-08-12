package main
import(
  "fmt"
)
func main(){
  var n, k int
  fmt.Scan(&n)
  fmt.Scan(&k)
  curr := 0
    queue := make([]int, n)
    for i := 0; i < n; i++{
      queue[i] = i+1
    }
    for len(queue) > 1{
     curr = (curr + k -1) % len(queue)
     queue = append(queue[:curr], queue[curr+1:]...)
     }
   fmt.Println(queue[0])
 }


