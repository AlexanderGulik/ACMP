package main
import(
  "fmt"
)

func main(){
  var n, k int
  fmt.Scan(&n, &k)
  for i :=n; ;i++{
    t := i
    for j := 0; j <n; j++ {
      if t % n != k {
        break
      }
      t = (t-k) - (t-k)/n
      if t <= 0 {
        break
      }
      if j == n-1{
        fmt.Println(i)
        return
      }
    }
  }
}
