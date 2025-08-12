package main
import(
"bufio"
	"fmt"
	"os"
	"strconv"
	"strings" 
)

func main(){
  reader := bufio.NewReader(os.Stdin)
  writer := bufio.NewWriter(os.Stdout)
  defer writer.Flush()
  line, _ := reader.ReadString('\n')
  fields := strings.Fields(line)
  m, _ := strconv.Atoi(fields[0])
  n, _ := strconv.Atoi(fields[1])
  a := make([][]int, m)
  for i := range a {
    a[i] = make([]int, n)
    line, _ = reader.ReadString('\n')
    row := strings.Fields(line)
    for j := 0; j < n; j++{
      a[i][j], _ = strconv.Atoi(row[j])
    }
  }

  aa := make([]int, m)
  bb := make([]int, n)
  for i := 0; i < m; i++{
    minVal :=a[i][0]
    for j := 1; j<n;j++{
      if minVal > a[i][j]{
        minVal = a[i][j]
      }
    }
    aa[i] =minVal
  }
  for i := 0; i< n; i++{
    maxVal = a[0][i]
    for j := 1; j < m; j++{
      if maxVal < a[j][i] {
        maxVal = a[j][i]
      }
    }
    bb[i] = maxVal
  }
  count := 0
  for i :=0; i <m; i++{
    for j := 0;j <n; j++{
        if aa[i] == bb[j]{
          count++
        }
    }
  }
  fmt.Fprintln(writer, count)
}
