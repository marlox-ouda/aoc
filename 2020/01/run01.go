package main

import (
  "fmt"
  "errors"
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

func main() {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  fmt.Printf("Input: %d\n", input_numbers)
  number_a, number_b, _ := find_pair(input_numbers)
  fmt.Printf("%d and %d\n", number_a, number_b)
}
