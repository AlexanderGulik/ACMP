package main

import (
  "fmt"
  "bufio"
  "os"
)

func main() {
   in := bufio.NewReader(os.Stdin)
   var nb, mb, nd, md int
   fmt.Fscan(in, &nb, &mb)
   base := make([]string, nb)
   for i := 0; i < nb; i++{
    fmt.Fscan(in, &base[i])
   }
   fmt.Fscan(in, &nd, &md)
   desert := make([]string, nd)
   for i := 0; i < nd; i++{
      fmt.Fscan(in, &desert[i])
   }
   count := 0
   for i := 0; i <=nd-nb; i++{
     for j := 0; j<= md-mb; j++{
        valid := true
        for x := 0; x<nb; x++{
          for y :=0; y<mb; y++{
            if base[x][y] == '#' {
              if desert[i+x][j+y] != '#' {
                valid = false
                break
              }
            }
          }
          if !valid {
            break
          }
        }
        if valid {
          count++
        }
     }
   }
   fmt.Println(count)
}
