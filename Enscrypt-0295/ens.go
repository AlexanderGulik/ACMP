package main

import(
  "fmt"
)

func main(){
  var n1, n2 string
  fmt.Scan(&n1, &n2)
  for shift := 0; shift < 26; shift++{
    desc := descypt(n1, shift)
    if check(desc, n2) {
     fmt.Println(desc)
     return
    }
  }
  fmt.Println("IMPOSSIBLE")
  return
}
func descypt(s string, shift int) string {
  result := make([]byte, len(s))
  for i := 0; i < len(s); i++{
    c := s[i] - byte(shift)
    if c < 'A' {
      c += 26
    }
    result[i] = c
  }
  return string(result)
}
func check(s1, s2 string) bool {
  if  len(s2) == 0 || len(s1) < len(s2){
    return false
  }
  for i := 0; i <= len(s1)-len(s2); i++{
    if s1[i:i+len(s2)] == s2{
      return true
    }
  }
  return false
}
