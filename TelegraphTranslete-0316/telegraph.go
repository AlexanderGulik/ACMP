package main
import (
  "fmt"
)

func main(){
  var n int
  fmt.Scan(&n)
  maxSum := 0
  cost := 0
  low := 0
  high := n
  for low <= high {
    mid := (low+high)/2
    currentCost := 7 * ((mid+99)/100)
    if mid+currentCost <= n {
      maxSum = mid
      cost = currentCost
      low = mid+1
    } else {
      high = mid-1
    }
  }
 
  fmt.Printf("%d %d", maxSum, cost)

}

