package main

import (
  "fmt"
  "strings"
  "strconv"
  "bufio"
  "os"
)

func main(){
  scanner := bufio.NewScanner(os.Stdin)
  scanner.Scan()
  n, _ := strconv.Atoi(scanner.Text())
  shelf := make([]int, 0 )
  result := make([]int, 0)
  
  for i := 0; i < n; i++{
    if !scanner.Scan() {
			break
		}
    line := strings.TrimSpace(scanner.Text())
    currStep := strings.Split(line, " ")
    if len(currStep) == 2 {
      a, _:= strconv.Atoi(currStep[0])
      b, _:= strconv.Atoi(currStep[1])
      if a == 1 {
        shelf = append([]int{b}, shelf...)
      } else if a == 2 {
        shelf = append(shelf, b)
      }
    } else if len(currStep) == 1 {
      drop, _ := strconv.Atoi(currStep[0])
      if drop == 3 {
        if len(shelf) > 0{
          last := shelf[0]
        result = append(result, last)
        shelf = shelf[1:]
      }
      } else if drop == 4 {
        if len(shelf) > 0 {
          last := shelf[len(shelf)-1]
        result = append(result, last)
        shelf = shelf[:len(shelf)-1]
      }
      }
    }
  }
  for i:= 0; i < len(result); i++{
    fmt.Printf("%d ", result[i])
  }
}
