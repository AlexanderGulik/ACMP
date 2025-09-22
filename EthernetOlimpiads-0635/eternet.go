package main

import (
  "fmt"
  "sort"
)

type Team struct {
  id int
  s int
  t int
}

func main() {
  var n, m int
  fmt.Scan(&n, &m)
  basic := make([]Team, n)
  for i := 0; i < n; i++{
    fmt.Scan(&basic[i].id, &basic[i].s, &basic[i].t)
  }
  advanced := make([]Team, m)
  for i := 0; i < m; i++{
    fmt.Scan(&advanced[i].id, &advanced[i].s, &advanced[i].t)
  }

  nextAdvanced := make(map[int]bool)
  for _, team := range advanced {
    if team.s > 0 {
      nextAdvanced[team.id] = true
    }
  }
  if n > 0 {
    basicWithSolutions := make([]Team, 0, n)
    for _, team := range basic {
      if team.s > 0 {
        basicWithSolutions = append(basicWithSolutions, team)
      }
    }

    if len(basicWithSolutions) > 0 {
      sort.Slice(basicWithSolutions, func(i, j int) bool {
        if basicWithSolutions[i].s == basicWithSolutions[j].s {
          return basicWithSolutions[i].t < basicWithSolutions[j].t
        }
        return basicWithSolutions[i].s > basicWithSolutions[j].s
      })
      med := len(basicWithSolutions)
      var proh int
      if med > 1 {
        proh = basicWithSolutions[med/2-1].s
      } else {
        proh = 0
      }

      for i, team := range basicWithSolutions {
        if i < med {
          if team.s > proh || team.s == basicWithSolutions[0].s {
            nextAdvanced[team.id] = true
          }
        }
      }
      
    }
  }

  result :=  make([]int, 0, len(nextAdvanced))
  for id := range nextAdvanced {
    result = append(result, id)
  }
  sort.Ints(result)
  fmt.Println(len(result))
  for i, id := range result {
    if i > 0 {
      fmt.Print(" ")
    }
    fmt.Print( id)
  }
  fmt.Println()
}
