package main
import (
  "fmt"
  "strconv"
  
)

func main(){
  var a, b string
  fmt.Scan(&a, &b)
  x1 := int(a[0]-'a')
  y1, _ := strconv.Atoi(string(a[1]))
  y1--
  x2 := int(b[0]-'a')
  y2, _ := strconv.Atoi(string(b[1])) 
  y2--
  if checkForOne(x1, y1, x2, y2) {
    fmt.Println(1)
    return
  }

  if checkForTwo(x1, y1, x2, y2) {
    fmt.Println(2)
    return
  }
    fmt.Println("NO")
}

func checkForOne(x1, y1, x2, y2 int) bool {
  dx := abs(x1-x2)
  dy := abs(y1-y2)
  return (dx==2 && dy ==1) || (dx == 1 && dy ==2)
}

func abs(x int) int {
  if x < 0 {
    return -x
  }
  return x
}

func checkForTwo(x1, y1, x2, y2 int) bool {
  for i := 0; i < 8; i++ {
    for j := 0; j < 8; j++{
      if checkForOne(x1,y1,i,j) && checkForOne(i, j, x2, y2) {
        return true
      }
    }
  }
  return false
}
