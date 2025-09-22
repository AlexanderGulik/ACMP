package main

import (
    "bufio"
    "fmt"
    "os"
    "sort"
)

type Person struct {
    t, w    int64
    index   int
}
func max(a, b int64) int64 {
  if a > b {
    return a
  }
  return b
}

func main() {
    in := bufio.NewReader(os.Stdin)
    var n int
    fmt.Fscan(in, &n)
    
    people := make([]Person, n)
    for i := 0; i < n; i++ {
        fmt.Fscan(in, &people[i].t, &people[i].w)
        people[i].index = i
    }
    
     sort.Slice(people, func(i, j int) bool {
        return people[i].w < people[j].w
    })
  result := make([]int64, n)
  maxTime := int64(0)
    for i := 0; i < n; i++ {
      currTime := people[i].t * people[i].w
      maxTime = max(maxTime, currTime)
      result[people[i].index] = maxTime
    }

    for i := 0; i< n; i++{
      fmt.Println(result[i])
    }
}
