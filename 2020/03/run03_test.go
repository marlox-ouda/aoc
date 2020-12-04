package main

import "testing"

func TestIsTree(t *testing.T) {
  var givenResult bool
  givenResult = isTree(example_geology, 0, 0)
  if givenResult {
    t.Error("isTree(example_geology, 0, 0) give true, false expected")
  }
  givenResult = isTree(example_geology, 1, 3)
  if givenResult {
    t.Error("isTree(example_geology, 1, 3) give true, false expected")
  }
  givenResult = isTree(example_geology, 2, 6)
  if !givenResult {
    t.Error("isTree(example_geology, 2, 6) give false, true expected")
  }
  givenResult = isTree(example_geology, 3, 9)
  if givenResult {
    t.Error("isTree(example_geology, 1, 3) give true, false expected")
  }
  givenResult = isTree(example_geology, 4, 12)
  if !givenResult {
    t.Error("isTree(example_geology, 2, 6) give false, true expected")
  }
  givenResult = isTree(example_geology, 5, 15)
  if !givenResult {
    t.Error("isTree(example_geology, 2, 6) give false, true expected")
  }
}

func TestCountTreesOnDirection(t *testing.T) {
  var (
    given_result int
    expected_result int
  )
  given_result = countTreesOnDirection(example_geology, 1, 1)
  expected_result = 2
  if given_result != expected_result {
    t.Errorf("countTreesOnDirection(example_geology, 1, 1) give %v, %v expected", given_result, expected_result)
  }
  given_result = countTreesOnDirection(example_geology, 3, 1)
  expected_result = 7
  if given_result != expected_result {
    t.Errorf("countTreesOnDirection(example_geology, 3, 1) give %v, %v expected", given_result, expected_result)
  }
  given_result = countTreesOnDirection(example_geology, 5, 1)
  expected_result = 3
  if given_result != expected_result {
    t.Errorf("countTreesOnDirection(example_geology, 5, 1) give %v, %v expected", given_result, expected_result)
  }
  given_result = countTreesOnDirection(example_geology, 7, 1)
  expected_result = 4
  if given_result != expected_result {
    t.Errorf("countTreesOnDirection(example_geology, 7, 1) give %v, %v expected", given_result, expected_result)
  }
  given_result = countTreesOnDirection(example_geology, 1, 2)
  expected_result = 2
  if given_result != expected_result {
    t.Errorf("countTreesOnDirection(example_geology, 1, 2) give %v, %v expected", given_result, expected_result)
  }
}

func TestMultiply(t *testing.T) {
  var (
    given_result int
    expected_result int
  )
  given_result = multiply([]int{2, 7, 3, 4, 2})
  expected_result = 336
  if given_result != expected_result {
    t.Errorf("multiply: %v expected, %v given", expected_result, given_result)
  }
}

func TestReadInput(t *testing.T) {
  input_lines, err := readInput()
  if err != nil {
    t.Errorf("readInput: fail with %s error", err)
  }
  expected_first_line := ".....##.#.....#........#....##."
  expected_last_line := ".#..##.##.#......#....##..#...."
  expected_len := 323
  if input_lines[0] != expected_first_line {
    t.Errorf("readInput: Not expected first line, %s expected, %s given",
             expected_first_line, input_lines[0])
  }
  given_len := len(input_lines)
  if given_len != expected_len {
    t.Errorf("readInput: Not expected slice len, %v expected, %v given",
             expected_len, given_len)
  }
  if input_lines[given_len-1] != expected_last_line {
    t.Errorf("readInput: Not expected last line, %s expected, %s given",
             expected_last_line, input_lines[given_len-1])
  }
}
