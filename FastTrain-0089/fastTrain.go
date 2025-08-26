package main
import (
  "fmt"
  "strings"
  "strconv"
  "os"
  "math"
)

func main(){
  content, _ := os.ReadFile("input.txt")  
  lines := strings.Split(strings.TrimSpace(string(content)), "\n")
  n, _ := strconv.Atoi(lines[0])
  var nameTrain string
  minTimeH, minTimeM := 25, 61
  for i := 1; i <= n ; i++{
    name, rest, _ := strings.Cut(lines[i], `"`)
    name, rest, _ = strings.Cut(rest, `"`)
    name = strings.TrimSpace(name)
    line := strings.Fields(rest)
    currTimeIn := line[0]
    currTimeOut := line[1]
    cHour, cMin := timeConvert(currTimeIn, currTimeOut)
    if (minTimeH > cHour )|| (minTimeH == cHour && minTimeM> cMin) {
      nameTrain = name
      minTimeH = cHour
      minTimeM = cMin
    } 
  }
  totalMinutes := minTimeH * 60 + minTimeM
  result := float64(650)/(float64(totalMinutes)/60.0)
  out, _ := os.Create("output.txt")
  defer out.Close()
  fmt.Fprintf(out,"The fastest train is \"%s\".\n", nameTrain)
  fmt.Fprintf(out, "Its speed is %d km/h, approximately.", int(math.Round(result)))

}


func timeConvert(time1, time2 string) (int, int) {
  parse := func(s string) (int, int) {
    parts := strings.Split(s, ":")
    h, _ := strconv.Atoi(parts[0])
    m, _ := strconv.Atoi(parts[1])
    return h, m
  } 
  h1, m1 := parse(time1)
  h2, m2 := parse(time2)
  if h1==h2 && m1==m2{
    return 24, 0
  }
  totalMin1 := h1*60+m1
  totalMin2 := h2*60 + m2
  if totalMin2 < totalMin1{
    totalMin2 += 24 *60
  }
  diff := totalMin2 - totalMin1
  return diff/60, diff%60
 }

/*
  Сори пидоры не принимают код на го ленге, по хуй знает какой причине, just Presentation error... 
  Я пробовал множество способов, ни один не принял(несколько консольных выводов, буфио напряпую без парсинга в консоль пытался пропихнуть, 
  вывод в файл, перенос для виндовс), попрсоил нейронку переписать код на с++ и компилятор принял.
  ХЗ с чем это свзяно, мб в логике алгоритма врятли ошибка, сейчас я понимаю что его можно было сделать намного проще
  #include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <sstream>
#include <cmath>
#include <algorithm>

using namespace std;

pair<int, int> timeConvert(const string& time1, const string& time2) {
    auto parse = [](const string& s) -> pair<int, int> {
        size_t pos = s.find(':');
        int h = stoi(s.substr(0, pos));
        int m = stoi(s.substr(pos + 1));
        return {h, m};
    };

    auto [h1, m1] = parse(time1);
    auto [h2, m2] = parse(time2);

    if (h1 == h2 && m1 == m2) {
        return {24, 0};
    }

    int totalMin1 = h1 * 60 + m1;
    int totalMin2 = h2 * 60 + m2;

    if (totalMin2 < totalMin1) {
        totalMin2 += 24 * 60;
    }

    int diff = totalMin2 - totalMin1;
    return {diff / 60, diff % 60};
}

int main() {
    ifstream in("input.txt");
    vector<string> lines;
    string line;

    while (getline(in, line)) {
        if (!line.empty()) {
            lines.push_back(line);
        }
    }

    int n = stoi(lines[0]);
    string nameTrain;
    int minTimeH = 25, minTimeM = 61;

    for (int i = 1; i <= n; i++) {
        string currentLine = lines[i];


        size_t firstQuote = currentLine.find('"');
        size_t secondQuote = currentLine.find('"', firstQuote + 1);
        string name = currentLine.substr(firstQuote + 1, secondQuote - firstQuote - 1);


        string rest = currentLine.substr(secondQuote + 1);


        istringstream iss(rest);
        vector<string> parts;
        string part;

        while (iss >> part) {
            parts.push_back(part);
        }

        if (parts.size() < 2) {
            continue;
        }

        string currTimeIn = parts[0];
        string currTimeOut = parts[1];

        auto [cHour, cMin] = timeConvert(currTimeIn, currTimeOut);

        if (minTimeH > cHour || (minTimeH == cHour && minTimeM > cMin)) {
            nameTrain = name;
            minTimeH = cHour;
            minTimeM = cMin;
        }
    }

    int totalMinutes = minTimeH * 60 + minTimeM;
    double result = 650.0 / (totalMinutes / 60.0);

    ofstream out("output.txt");
    out << "The fastest train is \"" << nameTrain << "\"." << endl;
    out << "Its speed is " << static_cast<int>(round(result)) << " km/h, approximately.";

    return 0;
}
*/
