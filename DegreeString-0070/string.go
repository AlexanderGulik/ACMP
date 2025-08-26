package main
import (
  "fmt"
  "strings"
)

func main(){
  var s string
  var n int
  fmt.Scan(&s)
  fmt.Scan(&n)

  if n > 0 {
      if n == 1 {
            outputLimited(s)
            return
        }
    totalNeeden := len(s) * n
    if totalNeeden > 1023 {
      output := make([]byte, 0, 1023)
      for len(output) < 1023 {
        for i := 0; i < len(s) && len(output) < 1023; i++{
          output = append(output, s[i])
        }
      }
    fmt.Println(string(output))
    return
    } else{
    var result strings.Builder
    result.Grow(len(s)*n)
    for i := 0; i < n; i++{
      result.WriteString(s)
    }
    outputLimited(result.String())
    return
  }
  } else {
    subString := findMinimalSubstring(s, -n)
       if "" == subString {
      fmt.Println("NO SOLUTION")
      return
    } else {
      outputLimited(subString)
      return
    }

  }
}

func findMinimalSubstring(s string, rep int ) string {
  if len(s) == 0 || rep == 0 {
    return ""
  }
  if rep == 1 {
    return s
  }
  if len(s) %rep != 0 {
    return ""
  }
  lenght := len(s) / rep
  subString := s[:lenght]
  sBytes := []byte(s)
  subBytes := []byte(subString)

  for i := 1; i < rep; i++{
    start := i * lenght
    for j := 0; j < lenght; j++ {
    if sBytes[start+j] != subBytes[j] {
      return ""
    }
    } 
  }
  return subString
}
func outputLimited(str string) {
    if len(str) > 1023 {
        fmt.Println(str[:1023])
    } else {
        fmt.Println(str)
    }
}
/* бля я ебал реп, хули мое решениене проходит по времени конкатинация сосет
   ЕБАЛ РЕП, СУКИ КАКАЯ НАХУЮ ОШИБКА, ПЕРЕПИСЫВАЮ ФУНКЦИЮ, но здесь оставлю на память
   func findMinimalSubstring(s string, rep int ) string {
  if len(s) == 0 || rep == 0 {
    return ""
  }
  if len(s) %rep != 0 {
    return ""
  }
  lenght := len(s) / rep
  subString := s[:lenght]
  for i := 0; i < len(s); i+= lenght{
    end := i + lenght
    if s[i:end] != subString {
      return ""
    }
  }
  return subString
}
прошлый вариант который не подошел, но функцию которую я написал и мои мысли при решении данной задачи хочу сохранить
package main
import (
  "fmt"
)

func main(){
  var s string
  var n int
  fmt.Scan(&s)
  fmt.Scan(&n)
  var result string
  if n > 0 {
    for i := 0; i < n; i++{
      result += s
    }
    fmt.Println(result)
    return
  } else {
    subString := findMinimalSubstring(s)
    var currString string
    for i:= 0; i < -n; i++{
      currString += subString
    }
    if s != currString {
      fmt.Println("NO SOLUTION")
      return
    } else {
      fmt.Println(subString)
    }
  /*как мне можно найти подстроку, в строке которая повротяется множество раз? или вообще не повторяется
  сначала мне нужно найти подстроку:
  для этого мне нужно написать функцию, пока не найден повторение, но каким же образом мне искать ее
  давай обдумаем случаи которые могут быть
  подстрока может состоять из ffqffq как мне определить подстроку ffq, ну скорее всего нужно перемещаться по строке с одним символом, пока
  я не найду максимальную длину подстроки, то есть алгоритм будет такой
  создаем переменную для хранения подстроки
  берем первый символ, сравниваем со следующим,

  если он повторяется, означает что наша подстрока состоит из одного символа, дальше сравниваем до конца строки,
  если он соответствует всей строке
  возвращаем единственный символ,
  иначе добавляем в подстроку символ и сравниваем со всет строкой
  после мне нужно понять что строка стостоит только из найденных подстрок подстрок
//
  
  }
}

func findMinimalSubstring(s string ) string {
  if len(s) == 0 {
    return ""
  }
  for length := 1; length <= len(s); length++{
    if len(s) % length != 0 {
      continue
    }
    subString := s[:length]
    valid := true
    for i := 0; i < len(s); i+= length{
      end := i + length
      if end > len(s) {
        end = len(s)
      }
      if s[i:end] != subString {
        valid = false
        break
      }
    }
    if valid {
      return subString
    }
  }
  return s
}

       
  */ 
