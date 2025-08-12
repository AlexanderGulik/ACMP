package main
import (
  "fmt"
  "bufio"
  "os"
  "strconv"
  "strings"
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()

  line, _ := reader.ReadString('\n')
  count := make([]int, 201)
  
  line, _ = reader.ReadString('\n')
  nums := strings.Fields(line)
    
    for _, s := range nums {
      temp, _ := strconv.Atoi(s)
      count[temp+100]++
    }

  for num := -100; num <= 100; num++{
  for i := 0 ; i < count[num+100]; i++ {
    fmt.Fprint(writer, num, " ")
  } 
}

}


