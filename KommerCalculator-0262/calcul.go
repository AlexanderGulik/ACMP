package main

import (
  "fmt"
  "sort"
)

func main() {
  var n int
  fmt.Scan(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
  }
  cost := 0

  for i:= 0; i < n-1; i++{ 
    sort.Ints(nums[i:])
    nums[i+1] += nums[i]
    cost += nums[i+1]
  }
  result := float64(cost) * 0.05
  fmt.Printf("%.2f",result)
}


/*#include <fstream>

#include <set>

#include <iomanip>

using namespace std;

int main() {

    ifstream ifst("input.txt");

    ofstream ofst("output.txt");

    multiset<long> numbers;

    size_t n;

    ifst >> n;

    for (size_t i = 0; i < n; ++i) {

        long number;

        ifst >> number;

        numbers.insert(number);

    }

    double cost = 0;

    while (numbers.size() > 1) {

        auto first = numbers.begin(),

            second = next(first),

            third = next(second);

        long sum = *first + *second;

        numbers.erase(first, third);

        cost += sum;

        numbers.insert(sum);

    }

    ofst << setprecision(2) << std::fixed << cost * 0.05;

}*/








/*
 func main() {
  var n int
  fmt.Scan(&n)
  nums := make([]int, n)
  for i := 0; i < n; i++{
      fmt.Scan(&nums[i])
  }
  sort.Ints(nums)
  cost := 0.0
  for len(nums) >1 {
    sum := nums[0] + nums[1]
    cost += float64(sum)
    nums = nums[2:]
    idx := sort.SearchInts(nums, sum)
    nums = append(nums, 0)
    copy(nums[idx+1:], nums[idx:])
    nums[idx] = sum
  }
  result := cost * 0.05
  fmt.Printf("%.2f",result)
}

*/


















