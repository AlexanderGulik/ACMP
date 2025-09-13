package main

import (
  "fmt"
)

func main() {
  var n, k int
  fmt.Scan(&n, &k)
  k--
  graph := make([][]int, n)
  reverseGraph := make([][]int, n)

  for {
    var x, y int
    fmt.Scan(&x)
    if x == 0 {
      break
    }
    fmt.Scan(&y)  
    x--
    y--
    graph[x] = append(graph[x], y)
    reverseGraph[y] = append(reverseGraph[y], x)
  }
  solwer := make([]bool, n)
  queue := []int{k}
  solwer[k] = true
  for len(queue) > 0 {
    curr := queue[0]
    queue = queue[1:]
    for _, neig := range graph[curr] {
      if !solwer[neig] {
        solwer[neig] = true
        queue = append(queue, neig)
      }
    }
  }
  faster := make([]bool, n)
  queue = []int{k}
  faster[k] = true

  for len(queue) > 0 {
    curr := queue[0]
    queue = queue[1:]
    for _, neig := range reverseGraph[curr]{
      if !faster[neig] {
        faster[neig] = true
        queue = append(queue, neig)
      }
    }
  }
  count := 0
  for i := 0; i < n; i++{
    if solwer[i] || faster[i] {
      count++
    }
  }
  if count == n {
   fmt.Println("Yes")
  } else {
   fmt.Println("No")
  }
}
