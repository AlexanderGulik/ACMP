package main
import(
  "fmt"
)

func main() {
  var x1, y1, x2, y2 int
  fmt.Scan(&x1, &y1)
  fmt.Scan(&x2, &y2)
  dx := abs(x2-x1)
  dy := abs(y2-y1)
  if (dx+dy)%2 != 0 {
    fmt.Println(0)
    return
  }
  if dx == 0 && dy == 0 {
    fmt.Println(0)
    return
  }
  if dx == dy {
    fmt.Println(1)
    return
  }
  fmt.Println(2)
}

func abs(a int) int {
  if a < 0{
    return -a
  }
  return a
}
