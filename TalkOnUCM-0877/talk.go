//последняя поптыка, пытался оптимизировать свой код с помощью гпт и кода питона (Presentaion error 4 test идите нахуй)
package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	file, _ := os.Open("input.txt")
	defer file.Close()
	scanner := bufio.NewScanner(file)
	talk := make([]string, 0)
	
	scanner.Scan()
	firstLine := scanner.Text()
	parts := strings.Split(firstLine, " ")
	companionName := ""
	for i, part := range parts {
		if part == "signed" {
			companionName = strings.Join(parts[1:i], " ")
			break
		}
	}

	for scanner.Scan() {
		row := scanner.Text()
		
		if strings.Contains(row, " signed off") {
			break
		}
		
		if len(row) < 10 {
			continue 
		}
		text := row[10:]
		
		if len(text) > 0 && text[len(text)-1] == '\n' {
			text = text[:len(text)-1]
		}
		
		if len(text) > 0 {
			lastChar := text[len(text)-1]
			if lastChar != '!' && lastChar != '?' {
				if lastChar == '.' || lastChar == ',' {
					text = text[:len(text)-1]
				}
				text += ","
			}
		}
		
		talk = append(talk, text)
	}

	out, _ := os.Create("output.txt")
	defer out.Close()
	output := bufio.NewWriter(out)
	defer output.Flush()
	
	for i := 0; i < len(talk); i++ {
		var speaker string
		if i%2 == 0 {
			speaker = "Fedya"
		} else {
			speaker = companionName
		}
		fmt.Fprintf(output, "\"%s\" --- skazal %s.\n", talk[i], speaker)
	}
}

/* мой код
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
)

func main(){
  file, _ := os.Open("input.txt")
  defer file.Close()
  scanner := bufio.NewScanner(file)
  talk := make([]string, 0)
  scanner.Scan()
  firstLine := scanner.Text()
  parts := strings.Split(firstLine, " ")
  companionName := ""
  for i, part := range parts {
    if part == "signed" {
        companionName = strings.Join(parts[1:i], " ")
        break
    }
  }
    for scanner.Scan() {
    row := string(scanner.Text())
    text := row[10:]

    if strings.Contains(row, " signed off"){
      break
    }
    char := text[len(text)-1]
    if char == '.' {
      text = text[:len(text)-1]+ ","
    } else if char == '!' || char == '?'{

    } else {
      text += ","
    }
       talk = append(talk, text)
  }
  out, _ := os.Create("output.txt")
  defer out.Close()
  output := bufio.NewWriter(out)
  defer output.Flush()
  for i := 0; i < len(talk); i++{
    var speaker string
    if i%2==0 {
      speaker = "Fedya"
    } else {
      speaker = companionName
    }
      fmt.Fprintf(output, "\"%s\" --- skazal %s.\n", talk[i],speaker)
  }
}*/
/*
я жестко ебал реп и хз почему Presentaion error в моем коде на 4 тесте
def main():
    input_file = open("input.txt", "r")
    output_file = open("output.txt", "w")
    with open('input.txt') as myfile:
        count = sum(1 for line in myfile)
    print(count)
    line = input_file.readline().split()
    name = str(line[1])
    print(name)
    for i in range(1, count - 1):
        line = input_file.readline()
        a = str(line)
        a = a[10:]
        if a[len(a) - 1] =='\n':
            a = a[:-1]
        if a[len(a) - 1] != '!' and a[len(a) - 1] != '?':
            if a[len(a) - 1] == '.' or a[len(a) - 1] == ',':
                a = a[:-1]
            a += ','
        print(a)
        if i % 2 == 0:
            ansname = name
        else:
            ansname = "Fedya"
        output_file.write(str("\"") + str(a) + "\"" + " --- skazal " + str(ansname) + "." + "\n")
    #output_file.write(str(0) + "\n")


if __name__ == "__main__":
    main()

*/
