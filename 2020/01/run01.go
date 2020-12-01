package main

import (
  "errors"
  "fmt"
  "os"
)

// find_pair returns the two firsts elements whose sum equal 2020
// take int slice as input
// third returned value is error or nil
func find_pair(s []int) (int, int, error) {
  for index, number_a := range s {
    for _, number_b := range s[index:] {
      if number_a + number_b == 2020 {
        return number_a, number_b, nil
      }
    }
  }
  return 0, 0, errors.New("can't find pair whose sum give 2020")
}

// 
func fix_expense (s []int) (int, error) {
  number_a, number_b, err := find_pair(s)
  if err != nil {
    return 0, err
  }
  return number_a * number_b, nil
}

func example_one() {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  fmt.Printf("Input: %d\n", input_numbers)
  result, _ := fix_expense(input_numbers)
  fmt.Printf("Output: %d\n", result)
}

func main() {
  request := "all"
  if len(os.Args) > 1 {
    request = os.Args[1]
  }
  if request == "1" || request == "example1" || request == "all" {
    example_one()
  }
}
