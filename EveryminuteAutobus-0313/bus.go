package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  line, _ := reader.ReadString('\n')
  line, _ = reader.ReadString('\n')
  nums := strings.Fields(line)
  result := count(nums)
  fmt.Fprint(writer, result)
}

func count(nums []string) int {
  maps := make(map[string]int)
  count := 0
  for j, num := range nums{
    if i, found := maps[num]; found{
      curr := j-i
      if curr > count {
        count = curr
      }
    } 
      maps[num] = j
    
  }
  return count
    
}

/* первый вариант, который слишком медленный
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  line, _ := reader.ReadString('\n')
  line, _ = reader.ReadString('\n')
  nums := strings.Fields(line)
  result := count(nums)
  fmt.Fprint(writer, result)
}
// перемещаемсся по массиву o(n^2)
func count(nums []string) int {
  count := 1
  for i := 0; i < len(nums)-1; i++{
    curr := 1
    for j := i+1; j < len(nums); j++{
      if nums[i] == nums[j]{
        if count < curr {
              count = curr
        }
        break
      }
      curr++
    }
  }
  return count
    
}
*/
