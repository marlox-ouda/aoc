package main

import (
  "bufio"
  "errors"
  "fmt"
  "os"
  "strconv"
  "strings"
)

// checkPolicyOne if letter occurence in password is between min and max
func checkPolicyOne(min int, max int, letter byte, password string) bool {
  var count int
  count = strings.Count(password, string(letter))
  if min <= count && count <= max {
    return true
  }
  return false
}

// checkPolicyTwo if letter is present at only one given of both
// given position in the password
func checkPolicyTwo(pos_a int, pos_b int, letter byte, password string) bool {
  var (
    check_a bool
    check_b bool
  )
  check_a = (password[pos_a - 1] == letter)
  check_b = (password[pos_b - 1] == letter)
  return check_a != check_b
}

// convertLine
// read element from the given string line
func convertLine(line string) (int, int, byte, string, error) {
  var (
    number_a int
    number_b int
    letter byte
    password string
    err error
    parsed_items_nb int
  )
  parsed_items_nb, err = fmt.Fscanf(strings.NewReader(line), "%d-%d %c: %s",
                                   &number_a, &number_b, &letter, &password)
  if err == nil && parsed_items_nb < 4 {
    err = fmt.Errorf("Only able to parse %v/4 items on %s", parsed_items_nb, line)
  }
  return number_a, number_b, letter, password, err
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

func runOne(line string, verbose bool) bool {
  if verbose {
    fmt.Printf("Input: %s ->", line)
  }
  min, max, letter, password, err := convertLine(line)
  if err == nil {
    result := checkPolicyOne(min, max, letter, password)
    if verbose {
      fmt.Printf("Output: %t\n", result)
    }
    return result
  } else {
    fmt.Printf("Error: %s\n", err)
  }
  return false
 }

func exampleOne() {
  var (
    count_valids int
    line string
  )
  var lines = []string{
    "1-3 a: abcde",
    "1-3 b: cdefg",
    "2-9 c: ccccccccc",
  }
  for _, line = range lines {
    if runOne(line, true) {
      count_valids += 1
    }
  }
  fmt.Printf("Final result: %v\n", count_valids)
}

func mainOne() {
  _, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
  }
}

func runTwo(line string, verbose bool) bool {
  if verbose {
    fmt.Printf("Input: %s -> ", line)
  }
  min, max, letter, password, err := convertLine(line)
  if err == nil {
    result := checkPolicyTwo(min, max, letter, password)
    if verbose {
      fmt.Printf("Output: %t\n", result)
    }
    return result
  } else {
    fmt.Printf("Error: %s\n", err)
  }
  return false
 }

func exampleTwo() {
  var (
    count_valids int
    line string
  )
  var lines = []string{
    "1-3 a: abcde",
    "1-3 b: cdefg",
    "2-9 c: ccccccccc",
  }
  for _, line = range lines {
    if runTwo(line, true) {
      count_valids += 1
    }
  }
  fmt.Printf("Final result: %v\n", count_valids)
}

func mainTwo() {
  _, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
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
