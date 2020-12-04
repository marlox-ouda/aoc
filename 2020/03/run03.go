package main

import (
  "bufio"
  "fmt"
  "os"
  "strings"
)

var example_geology = []string{
  "..##.......",
  "#...#...#..",
  ".#....#..#.",
  "..#.#...#.#",
  ".#...##..#.",
  "..#.##.....",
  ".#.#.#....#",
  ".#........#",
  "#.##...#...",
  "#...##....#",
  ".#..#...#.#",
}

// isTree return whether there is a three or no at the given position
// position is given 0 based and take top left corner as (0, 0) position
func isTree(grid []string, vertical_offset int, horizontal_offset int) bool {
  line_len := len(grid[0])
  if vertical_offset > len(grid) {
    return false
  }
  horizontal_offset = horizontal_offset % line_len
  if grid[vertical_offset][horizontal_offset] == '#' {
    return true
  }
  return false
}


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

// Read input1.txt file and return a slice of int
// considering each line is just containing one int
func readInput( ) ([]string, error) {
  input_fd, err := os.Open("./input3.txt")
  if err != nil {
    return nil, err
  }
  defer input_fd.Close()
  scanner := bufio.NewScanner(input_fd)
  var lines []string
  var line string
  for scanner.Scan() {
    line = scanner.Text()
    lines = append(lines, line)
  }
  return lines, nil
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
  fmt.Printf("ExampleOne: Final result: %v\n", count_valids)
}

func mainOne() {
  var (
    count_valids int
    line string
  )
  lines, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    for _, line = range lines {
      if runOne(line, false) {
        count_valids += 1
      }
    }
    fmt.Printf("MainOne: Final result: %v\n", count_valids)
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
    fmt.Printf("MainOne: Error: %s\n", err)
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
  fmt.Printf("ExampleTwo: Final result: %v\n", count_valids)
}

func mainTwo() {
  var (
    count_valids int
    line string
  )
  lines, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    for _, line = range lines {
      if runTwo(line, false) {
        count_valids += 1
      }
    }
    fmt.Printf("MainTwo: Final result: %v\n", count_valids)
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
