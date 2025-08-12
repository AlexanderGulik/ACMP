package main
import(
  "fmt"
  "sort"
  "strings"
)

func main(){
  var n, m int
  fmt.Scan(&n)
  friends := make(map[string]bool)
  sortName := []string{}
  for i:=0; i< n; i++{
    var curr string
    fmt.Scan(&curr)
    friends[curr] = true
    sortName = append(sortName, curr)
  }
 
  mutalFriend := []string{}
  AlsoFriendOf := []string{}
  fmt.Scan(&m)
  for i:= 0; i < m; i++{
    var curr string
    fmt.Scan(&curr)
    if friends[curr] {
      mutalFriend = append(mutalFriend, curr)
    } else {
      AlsoFriendOf = append(AlsoFriendOf, curr)
    }
  }
  sort.Strings(sortName)
  sort.Strings(mutalFriend)
  sort.Strings(AlsoFriendOf)
  resultFirends := strings.Join(sortName, ", ")
  resultMutual := strings.Join(mutalFriend, ", ")
  resultAlso := strings.Join(AlsoFriendOf, ", ")
  fmt.Println("Friends:", resultFirends)
  fmt.Println("Mutual Friends:", resultMutual)
  fmt.Println("Also Friend of:", resultAlso)

}

