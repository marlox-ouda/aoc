package main

import (
  "bufio"
  "fmt"
  "os"
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
  if vertical_offset >= len(grid) {
    return false
  }
  horizontal_offset = horizontal_offset % line_len
  if grid[vertical_offset][horizontal_offset] == '#' {
    return true
  }
  return false
}

// count encounter trees from the top to the bottom of the grid
func countTreesOnDirection(grid []string, horizontal_step int, vertical_step int) int {
  var (
    current_vertical_offset int
    current_horizontal_offset int
    currenly_encountered_trees int
  )
  if isTree(grid, current_vertical_offset, current_horizontal_offset) {
      currenly_encountered_trees += 1
  }
  for current_vertical_offset < len(grid) {
    current_vertical_offset += vertical_step
    current_horizontal_offset += horizontal_step
    if isTree(grid, current_vertical_offset, current_horizontal_offset) {
      currenly_encountered_trees += 1
    }
  }
  return currenly_encountered_trees
}

// multiply every component of the slice between them and return the value
func multiply(values []int) int {
  var (
    total int
    current_value int
  )
  total = 1
  for _, current_value = range values {
    total *= current_value
  }
  return total
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

func exampleOne() {
  var trees_number int
  trees_number = countTreesOnDirection(example_geology, 3, 1)
  fmt.Printf("ExampleOne: %v\n", trees_number)
}

func mainOne() {
  var trees_number int
  lines, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    trees_number = countTreesOnDirection(lines, 3, 1)
    fmt.Printf("MainOne: %v\n", trees_number)
  }
}

func runTwo(grid []string) int {
  var (
    trees_number int
    trees_numbers []int
    direction []int
    horizontal_step int
    vertical_step int
  )
  directions := [][]int{
    []int{1, 1},
    []int{3, 1},
    []int{5, 1},
    []int{7, 1},
    []int{1, 2},
  }

  for _, direction = range directions {
    horizontal_step = direction[0]
    vertical_step = direction[1]
    fmt.Printf("Direction: (%v, %v) -> ", horizontal_step, vertical_step)
    trees_number = countTreesOnDirection(grid, horizontal_step, vertical_step)
    fmt.Printf("%v tree(s)\n", trees_number)
    trees_numbers = append(trees_numbers, trees_number)
  }
  return multiply(trees_numbers)
 }

func exampleTwo() {
  fmt.Printf("ExampleTwo: %v\n", runTwo(example_geology))
}

func mainTwo() {
  lines, err := readInput()
  if err != nil {
    fmt.Println(err)
  } else {
    fmt.Printf("MainTwo: %v\n", runTwo(lines))
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
