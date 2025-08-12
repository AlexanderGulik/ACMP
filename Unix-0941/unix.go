package main
import(
  "fmt"
  "strconv"
)

func main(){
  var a, b string
  fmt.Scan(&a, &b)
  a1, _ := strconv.ParseInt(a, 3, 64)

  a2, _ := strconv.ParseInt(b, 3, 64)
  fmt.Println(a1+a2)
}
