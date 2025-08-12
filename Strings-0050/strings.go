package main
import(
  "fmt"
)

func main(){
  var a, b string
  fmt.Scan(&a)
  fmt.Scan(&b)
  lenB := len(b)
  if lenB == 0 || len(a) < lenB {
    fmt.Println(0)
    return
  }
  shifts := make(map[string]struct{})
  for k := 0; k < lenB; k++{
    shifted := b[k:] + b[:k]
    shifts[shifted] = struct{}{}
  }
  count := 0
  lenA := len(a)
  for i:= 0; i <= lenA - lenB; i++{
    sub := a[i : i+lenB]
    if _, exists := shifts[sub]; exists {
      count++
    }
  }
  fmt.Println(count)
}
