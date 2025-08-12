package main
import(
  "fmt"
  "sort"
)

func main(){
  var n, k int
  fmt.Scan(&n, &k)
  numbers := make([]int, n) 
  for i := 0; i < n; i++{
    var curr int
    fmt.Scan(&curr)
    numbers[i] = curr
  }
  sort.Sort(sort.Reverse(sort.IntSlice(numbers)))
  result := 0
  usesMen := 0
  currMen := 0
  for i := 0; i < n; i++{

    if currMen == k{
      usesMen++
      currMen = 0
    }
    result += numbers[i] * (usesMen+1)
    currMen++
   }
  fmt.Println(result)
}


