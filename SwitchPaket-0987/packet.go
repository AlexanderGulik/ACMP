package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.Open("input.txt")
	defer input.Close()
	output, _ := os.Create("output.txt")
	defer output.Close()

	scanner := bufio.NewScanner(input)
	scanner.Scan()
	firstLine := scanner.Text()
	nameIp := strings.Split(firstLine, " ")[1]
	
	packet := []int{}
	
	for i := 0; i < 4; i++ {
		scanner.Scan()
		row := scanner.Text()
		if row == "Time out" {
			continue
		}
		if strings.Contains(row, "Time=") {
			parts := strings.Split(row, "Time=")
			timeStr := strings.TrimSpace(parts[1])
			time, _ := strconv.Atoi(timeStr)
			packet = append(packet, time)
		}
	}

	lost := 4 - len(packet)
	lossPercent := lost * 25

	if len(packet) == 0 {
		fmt.Fprintf(output, "Ping statistics for %s:\n", nameIp)
		fmt.Fprintf(output, "Packets: Sent = 4 Received = 0 Lost = 4 (100%% loss)")
	} else {
		minP := packet[0]
		maxP := packet[0]
		sum := 0
		
		for _, time := range packet {
			if time < minP {
				minP = time
			}
			if time > maxP {
				maxP = time
			}
			sum += time
		}
		avg := (sum + len(packet)/2) / len(packet)
		
		fmt.Fprintf(output, "Ping statistics for %s:\n", nameIp)
		fmt.Fprintf(output, "Packets: Sent = 4 Received = %d Lost = %d (%d%% loss)\n", len(packet), lost, lossPercent)
		fmt.Fprintf(output, "Approximate round trip times:\n")
		fmt.Fprintf(output, "Minimum = %d Maximum = %d Average = %d", minP, maxP, avg)
	}
}
/* first try
package main

import (
  "fmt"
  "bufio"
  "os"
  "strings"
  "strconv"
)

func main() {
  input, _ := os.Open("input.txt")
  defer input.Close()
  scanner := bufio.NewScanner(input)
  var nameIp string
  scanner.Scan()
  currString := scanner.Text()
  nameIp = strings.Split(currString, " ")[1]
  packet := []int{}
  for i := 0; i < 4; i++ {
    if scanner.Scan(){
      row := scanner.Text()
    if !strings.Contains(row, "Time out") {
       s := strings.Split(row, " ")
       lastPart := s[len(s)-1]
       if strings.Contains(lastPart, "Time=") {
       indx := strings.Split(lastPart, "=")[1]
       time, _ := strconv.Atoi(indx)
        packet = append(packet, time)
    }
    }
  }
}
  if len(packet) == 0 {
    fmt.Printf("Ping statistics for %d:\n", nameIp)
    fmt.Printf("Packets: Sent = 4 Received = 0 Lost = 4 (100% loss)")
  } else {
    minP := min(packet)
    maxP := max(packet)
    averageP := average(packet)
    lossP := 100-((len(packet) * 100)/4)
    fmt.Printf("Ping statistics for %s:\n", nameIp)
    fmt.Printf("Packets: Sent = 4 Received = %d Lost = %d (%d%% loss)\n", len(packet), 4-len(packet), lossP)
    fmt.Printf("Approximate round trip times:\n")
    fmt.Printf("Minimum = %d Maximum = %d Average = %d", minP, maxP, averageP)
  }
  
}


func average(n []int) int {
  result := 0
  for i:=0; i< len(n); i++{
    result += n[i]
  }
  if len(n) > 0 {
   return result/ len(n)
  }
  return 0
}

func min(n []int) int{
  result := n[0]
  for i := 0; i < len(n); i++{
    if n[i] < result {
      result = n[i]
    }
  }
  return result
}

func max(n []int)int{
  result := n[0]
  for i := 0; i < len(n); i++{
    if n[i] > result{
      result = n[i]
    }
  }
  return result
}*/
