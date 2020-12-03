package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
)

// checkPolicyOne if letter is present at only one given of both
// given position in the password
func checkPolicyOne(pos_a int, pos_b int, letter byte, password string) bool {
  return true
}

// checkPolicyTwo if letter is present at only one given of both
// given position in the password
func checkPolicyTwo(pos_a int, pos_b int, letter byte, password string) bool {
  return true
}

// findPair returns the two first elements whose sum equal 2020
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

// findThree returns the three first elements whose sum equal 2020
// take int slice as input
// fourth returned value is error or nil
func findThree(s []int) (int, int, int, error) {
  for outer_index, number_a := range s {
    for inner_index, number_b := range s[outer_index:] {
      for _, number_c := range s[inner_index:] {
        if number_a + number_b + number_c == 2020 {
          return number_a, number_b, number_c, nil
        }
      }
    }
  }
  return 0, 0, 0, errors.New("can't find three numbers whose sum give 2020")
}

// fixExpense according Elves logic
func fixExpense (numbers_to_multiply ... int) int {
  total := 1
  for _, number := range numbers_to_multiply {
    total *= number
  }
  return total
}

// Read input1.txt file and return a slice of int
// considering each line is just containing one int
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

func runOne(input_numbers []int) {
  number_a, number_b, err := findPair(input_numbers)
  if err != nil {
    fmt.Println(err)
  } else {
    fixed_expense := fixExpense(number_a, number_b)
    fmt.Printf("Output: %d\n", fixed_expense)
  }
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
  runOne(input_numbers)
}

func mainOne() {
  input_numbers, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    runOne(input_numbers)
  }
}

func runTwo(input_numbers []int) {
  number_a, number_b, number_c, err := findThree(input_numbers)
  if err != nil {
    fmt.Println(err)
  } else {
    fixed_expense := fixExpense(number_a, number_b, number_c)
    fmt.Printf("Output: %d\n", fixed_expense)
  }
}

func exampleTwo() {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  fmt.Printf("Input: %d\n", input_numbers)
  runTwo(input_numbers)
}

func mainTwo() {
  input_numbers, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    runTwo(input_numbers)
  }
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
    mainOne()
  }
  if request == "2" || request == "example2" || request == "all" {
    exampleTwo()
  }
  if request == "2" || request == "input2" || request == "all" {
    mainTwo()
  }
}
