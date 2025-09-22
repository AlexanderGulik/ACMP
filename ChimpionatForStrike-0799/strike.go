package main

import (
  "fmt"
  "sort"
)

func main() {
  var n int
  fmt.Scan(&n)
  strike := make([]int, n)
  for i := 0; i < n; i++{
    fmt.Scan(&strike[i])
  }
  winner := 0
  for i := 0; i < n; i++ {
        if strike[i] > winner {
            winner = strike[i]
        }
    }

  hasWinnerBefore := make([]bool, n)
  found := false
  for i := 0; i < n; i++{
    hasWinnerBefore[i] = found
    if strike[i] == winner{
      found = true
    }
  }
  bestPlace := 0
  sortedScores := make([]int, n)
  copy(sortedScores, strike)
  sort.Sort(sort.Reverse(sort.IntSlice(sortedScores)))
  for i := 0; i< n-1; i++ {
    if strike[i] % 10 == 5 && hasWinnerBefore[i] && strike[i] > strike[i+1]{
        
      place := findsPlace(sortedScores, strike[i]) 
        if bestPlace == 0 || place < bestPlace {
            bestPlace = place
        }
    }
  }
      fmt.Println(bestPlace)
}

func findsPlace(sortedScores []int, target int) int {
  left, right := 0, len(sortedScores)-1
  result := -1
  for left <= right {
    mid := (left + right) / 2
    if sortedScores[mid] == target {
      result = mid
      right = mid-1
    } else if sortedScores[mid] > target {
      left = mid +1
    } else {
      right = mid -1
    }
  }
  if result != -1 {
    return result +1
  }
  return left + 1
}


