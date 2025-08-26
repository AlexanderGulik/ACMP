package main
import (
  "fmt"
  "sort"
)

func main(){
  var n int
  fmt.Scan(&n)
  nums := make([]int, n)
  orig := make([]int, n)

  for i:=0; i < n; i++{
    fmt.Scan(&nums[i])
    orig[i] = nums[i]
  }
 
  sort.Ints(nums)
  element := maxCountElement(nums)
  result := swapArr(orig, element)
  for i := 0; i < n; i++{
    fmt.Printf("%d ", result[i])
  }

}
/* первое что мне пришло в голову))) но вовремя осознал другой, более простой метод

func swapArr(nums []int, elem int){
   curr := 0
  for i := 0; i < len(nums); i++{
    if i+1 < len(nums){
      if nums[i] == elem{
        if nums[i] == nums[i+1] {
      for j := i+1; j < len(nums); j++{
         if nums[j] == elem{
          continue
        } else {
          curr = j
          break 
        }
      }
      nums[i], nums[curr] = nums[curr], nums[i]
    } else {
      nums[i], nums[i+1] = nums[i+1], nums[i]
    }
   }
  }
  }
}
*/
func swapArr(nums []int, elem int)[]int{
  var nonElems []int
  var elems []int
  for _, num := range nums{
    if num == elem {
      elems = append(elems, num)
    } else {
      nonElems = append(nonElems, num)
    }
  }
  return append(nonElems, elems...)
  }


/* можно было бы использховать хеш таблицу, вот пример кода c хеш таблицей. 
func maxCountElement(nums []int) int {
    counts := make(map[int]int)
    maxCount := 0
    result := nums[0] 
    for _, num := range nums {
        counts[num]++
        if counts[num] > maxCount {
            maxCount = counts[num]
            result = num
        } else if counts[num] == maxCount && num < result {
            result = num
        }
    }

    return result
}
Да эта функция быстрее, но мне впревую очередь пришла идея с сортировкой*/
func maxCountElement(nums []int) int  {
  result := nums[0]
  MaxCount := 1
  currentCount := 1
  for i := 1; i < len(nums); i++{
      if nums[i] == nums[i-1] {
        currentCount++
      } else {
        currentCount = 1
      }
      if( currentCount > MaxCount) || (MaxCount == currentCount && nums[i] < result ) {
         result = nums[i]
        MaxCount = currentCount
    
      }
        }
  return result
}

