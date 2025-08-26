package main
import (
  "fmt"
)

func main(){
  var queen, rock, horse string
  fmt.Scan(&queen, &rock, &horse)
  //первая мысль, что я могу сделать. 
  // 1. Я могу построить модель и посчитать на матрице клетки, котоыре участвуют
  // Я попробую сначала найти другой способ, если не получится - реализую первый
  // первый вариант не вышел, я понял что слишком сильно все запутал, первый гораздо проще

  queenX, queenY := parseChessPosition(queen)
  rockX,  rockY := parseChessPosition(rock)
  horseX, horseY := parseChessPosition(horse)
  var place [8][8]bool 
  for i := 0; i < 8; i++{
    for j := 0; j< 8; j++{
      place[i][j] = false
    }
  }
 
  queenPosition(&place, queenX, queenY)
  rockPosition(&place, rockX, rockY)
  horsePosition(&place, horseX, horseY)
  result := 0
  for i := 0; i < 8; i++{
    for j := 0; j< 8; j++{
      if place[i][j]{
        result++
      }
    }
  }

  fmt.Println(result-3) 
}
func parseChessPosition(pos string) (int, int) {
    if len(pos) != 2 {
        return 0, 0
    }
    letter := pos[0]
    words := map[string]int{
      "A": 0,
      "B": 1,
      "C": 2,
      "D": 3,
      "E": 4,
      "F": 5,
      "G": 6,
      "H": 7,
    }
    
    x, _:= words[string(letter)]
    digit := pos[1]

    y := int(digit - '1')  

    return x, y
}

func horsePosition(place *[8][8]bool, horseX, horseY int) {
  (*place)[horseX][horseY] = true
 if (horseX+1 < 8 && horseY-2 >=0)  { (*place)[horseX+1][horseY-2] = true } 
 if (horseX+1 < 8 && horseY+2 <8)  { (*place)[horseX+1][horseY+2] = true } 
 if (horseX-1 >= 0 && horseY+2 <8)  { (*place)[horseX-1][horseY+2] = true } 
 if (horseX-1 >= 0 && horseY-2 >=0)  { (*place)[horseX-1][horseY-2] = true } 
 
 if (horseX+2 < 8 && horseY-1 >=0)  { (*place)[horseX+2][horseY-1] = true } 
 if (horseX+2 < 8 && horseY+1 <8)  { (*place)[horseX+2][horseY+1] = true } 
 if (horseX-2 >= 0 && horseY-1 >=0)  { (*place)[horseX-2][horseY-1] = true } 
 if (horseX-2 >= 0 && horseY+1 <8)  { (*place)[horseX-2][horseY+1] = true }

}


func rockPosition(place *[8][8]bool, rockX, rockY int){
  (*place)[rockX][rockY] = true
  for i := 0; i < 8; i++{
    (*place)[i][rockY] = true
  }
  for i := 0; i < 8; i++{
    (*place)[rockX][i] = true
  }

}

func queenPosition(place *[8][8]bool, queenX, queenY int){
  (*place)[queenX][queenY] = true
  for i := 0; i < 8; i++{
   (*place)[i][queenY] = true 
   (*place)[queenX][i] = true
  }
    for i, j := queenX, queenY; i >= 0 && j >= 0; i, j = i-1, j-1 {
        (*place)[i][j] = true
    }
    for i, j := queenX, queenY; i < 8 && j < 8; i, j = i+1, j+1 {
        (*place)[i][j] = true
    }
    for i, j := queenX, queenY; i >= 0 && j < 8; i, j = i-1, j+1 {
        (*place)[i][j] = true
    }
    for i, j := queenX, queenY; i < 8 && j >= 0; i, j = i+1, j-1 {
        (*place)[i][j] = true
    }
 }

/* отхотды первого варианта:
func horsePosition(queenX, queenY, rockX, rockY, horseX, horseY int) int{
  result := 0
 if (horseX+1 <= 8 && horseY-2 >=1) &&  (horseX+1 != queenX && horseY-2 != queenY) &&  (horseX+1 != rockX && horseY-2 != rockY) { result++ } 
 if (horseX+1 <= 8 && horseY+2 <=8) &&  (horseX+1 != queenX && horseY+2 != queenY) &&  (horseX+1 != rockX && horseY+2 != rockY) { result++ } 
 if (horseX-1 >= 1 && horseY+2 <=8) &&  (horseX-1 != queenX && horseY+2 != queenY) &&  (horseX-1 != rockX && horseY+2 != rockY) { result++ } 
 if (horseX-1 >= 1 && horseY-2 >=1) &&  (horseX-1 != queenX && horseY-2 != queenY) &&  (horseX-1 != rockX && horseY-2 != rockY) { result++ } 
 
 if (horseX+2 <= 8 && horseY-1 >=1) &&  (horseX+2 != queenX && horseY-1 != queenY) &&  (horseX+2 != rockX && horseY-1 != rockY) { result++ } 
 if (horseX+2 <= 8 && horseY+1 <=8) &&  (horseX+2 != queenX && horseY+1 != queenY) &&  (horseX+2 != rockX && horseY+1 != rockY) { result++ } 
 if (horseX-2 >= 1 && horseY-1 >=1) &&  (horseX-2 != queenX && horseY-1 != queenY) &&  (horseX-2 != rockX && horseY-1 != rockY) { result++ } 
 if (horseX-2 >= 1 && horseY+1 <=8) &&  (horseX-2 != queenX && horseY+1 != queenY) &&  (horseX-2 != rockX && horseY+1 != rockY) { result++ } 
 return result
}


func rockPosition(queenX, queenY, rockX, rockY, horseX, horseY int) int{
  result := 14
    if queenX == rockX || queenY == rockY{
      result--
    }
    if rockX == horseX || rockY == horseY{
      result--
    }
    return result
}


func queenPosition(queenX, queenY, rockX, rockY, horseX, horseY int) int{
    result := 0
    if queenX == rockX || queenY == rockY{
      result--
    }
    if queenX == horseX || queenY == horseY{
      result--
    }
    fmt.Println(result)
    if queenX == 1 || queenY == 1 || queenX == 8 || queenY == 8 {
      result += 7+7+7
    } else {
      words := map[int]int{
       1: 8,
       2: 7,
       3: 6,
       4: 4,
       5: 5,
       6: 6,
       7: 7,
       8: 8,
      }
      value, _ := words[queenY]
      result = result + 7 + 7 + (7+(7- value)*2)
    }
    return result
}
*/
