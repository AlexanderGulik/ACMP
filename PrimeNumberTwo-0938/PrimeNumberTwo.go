package main
import(
  "os"
  "fmt"
  "strings"
  "strconv"
  "bufio"

)

func main() {
  content, _ := os.ReadFile("input.txt")
  line := strings.Split(strings.TrimSpace(string(content)), "\n")
  nums := strings.Split(strings.TrimSpace(line[1]), " ")
  maxCurrent := 0
  result := 0
    for i := 0; i <len(nums); i++{
      num, _ := strconv.Atoi(nums[i])
      current := primeNumber(num)
      if current > maxCurrent {
        maxCurrent = current
        result = num
      } else if current == maxCurrent && result > num {
        result = num
      }
    } 
  outFile, _ := os.Create("output.txt")
  defer outFile.Close()
  writer := bufio.NewWriter(outFile)
  fmt.Fprintln(writer, result)
  writer.Flush()
}

func primeNumber (num int) int {
  sum := 0
  if num % 2 == 0{
   sum++
   for num % 2 == 0 {
     num /=2
   }
  }

  for i := 3; i*i <= num; i+=2 {
    if num % i == 0{
      sum++
      for num % i ==0 {
        num /= i
      }
    }
  }
  if num > 1 {
    sum++
  }
  
 fmt.Println(num, sum) 
  return sum
}
