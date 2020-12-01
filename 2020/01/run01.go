package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
)

// find_pair returns the two firsts elements whose sum equal 2020
// take int slice as input
// third returned value is error or nil
func findPair(s []int) (int, int, error) {
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
func fixExpense (s []int) (int, error) {
  number_a, number_b, err := findPair(s)
  if err != nil {
    return 0, err
  }
  return number_a * number_b, nil
}

func readInput( ) ([]int, error) {
  input_fd, err := os.Open("./input1.txt")
  if err != nil {
    return nil, err
  }
  defer input_fd.Close()
  scanner := bufio.NewScanner(input_fd)
  var numbers_slice []int
  var casted_number int
  var line string
  for scanner.Scan() {
    line = scanner.Text()
    casted_number, err = strconv.Atoi(line)
    if err != nil {
      return nil, fmt.Errorf("can't cast %v to integer", line)
    }
    numbers_slice = append(numbers_slice, casted_number)
  }
  return numbers_slice, nil
}

func exampleOne() {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  fmt.Printf("Input: %d\n", input_numbers)
  result, _ := fixExpense(input_numbers)
  fmt.Printf("Output: %d\n", result)
}

func one() {
  readInput()
}

func main() {
  request := "all"
  if len(os.Args) > 1 {
    request = os.Args[1]
  }
  if request == "1" || request == "example1" || request == "all" {
    exampleOne()
  }
  if request == "1" || request == "input1" || request == "all" {
    one()
  }
}
