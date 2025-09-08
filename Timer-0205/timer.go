package main
import (
  "fmt"
  "strings"
)

func main() {
  var time, wait string
  fmt.Scan(&time)
  fmt.Scan(&wait)
  parts := strings.Split(time, ":")
  hours := parseInt(parts[0])
  minutes := parseInt(parts[1])
  seconds := parseInt(parts[2])
  totalSeconds := hours*3600 + minutes*60 + seconds
  intervals := strings.Split(wait, ":")
  var intervalsHours, intervalsMinutes, intervalsSeconds int
  switch len(intervals) {
    case 1:
      intervalsSeconds = parseInt(intervals[0])
    case 2:
      intervalsMinutes = parseInt(intervals[0])
      intervalsSeconds = parseInt(intervals[1])
    case 3:
      intervalsHours = parseInt(intervals[0])
      intervalsMinutes = parseInt(intervals[1])
      intervalsSeconds = parseInt(intervals[2])
  }
  intervalsTotalSeconds := intervalsHours*3600 + intervalsMinutes*60 +intervalsSeconds
  finalSeconds := totalSeconds + intervalsTotalSeconds
  days := finalSeconds / 86400
  remain := finalSeconds % 86400
  resultH := remain/3600
  remain %= 3600
  resultM := remain/60
  resultS := remain%60
  result := fmt.Sprintf("%02d:%02d:%02d", resultH, resultM, resultS)
  if days > 0 {
    result += fmt.Sprintf("+%d days", days)
  }
  fmt.Println(result)
}

func parseInt(s string) int {
	var result int
	fmt.Sscanf(s, "%d", &result)
	return result
}
