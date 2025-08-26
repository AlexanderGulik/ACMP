package main
import (
  "fmt"
)

func main (){
  var a, b, c string
  fmt.Scan(&a, &b, &c)
  max1 :=  max(a, b)
  max2 := max(b, c)
  fmt.Println( max(max1, max2))

}

func max(a, b string) string {
  if len(a) > len(b){
    return a
  } else if len(b) > len(a){
    return b
  }
  for i := 0; i < len(b); i++{
    if a[i] > b[i]{
      return a
    } else if a[i] < b[i]{
      return b
    }
  }
  return a
}
