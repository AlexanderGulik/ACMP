package main

import (
    "bufio"
    "fmt"
    "os"
)

func main(){
  in := bufio.NewReader(os.Stdin)
  out := bufio.NewWriter(os.Stdout)
  defer out.Flush()
  var n int
  fmt.Fscan(in, &n)
  houses := make([]int, n)
  for i := 0; i< n; i++{
    fmt.Fscan(in, &houses[i])
  }
  fmt.Fprintln(out, houses[n/2])
}

