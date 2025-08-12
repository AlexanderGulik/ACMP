package main
import (
  "fmt"
  )

func main(){
  n, s := 0,0
  fmt.Scan(&n)
  x := make([]int, n)
  y := make([]int, n)
  z1 := make([]bool, n)
  for i := range z1{
    z1[i] = true
  }
  for i := 0; i<n; i++{
    fmt.Scan(&x[i], &y[i])
  }
  for z := 0; z < n-1; z++{
    if z1[z]{
      for v := z+1; v < n; v++{
        if z1[v] && (x[z]*y[v] == x[v]*y[z]) && x[z] * x[v] >= 0 && y[z] * y[v] >= 0 {
          z1[v]= false
        }
      }
    }
  }
  for z := 0; z < n; z++{
    if z1[z]{
      s++
    }
  }
  fmt.Println(s)
}
