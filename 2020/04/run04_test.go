package main

import "testing"

func TestExtractPassportData(t *testing.T) {
  var (
    given_result *passport
    err error
  )
  given_result, err = extractPassportData("ecl:gry pid:860033327 eyr:2020 hcl:#fffffd\nbyr:1937 iyr:2017 cid:147 hgt:183cm")
  if err != nil {
    t.Errorf("extractPassport(<1>) raise error: %s", err)
  } else {
    if given_result.byr != "1937" {
        t.Errorf("extractPassport(<1>).byr, %s given, \"1937\" expected", given_result.byr)
    }
    if given_result.iyr != "2017" {
        t.Errorf("extractPassport(<1>).iyr, %s given, \"2017\" expected", given_result.iyr)
    }
    if given_result.eyr != "2020" {
        t.Errorf("extractPassport(<1>).eyr, %s given, \"2020\" expected", given_result.eyr)
    }
    if given_result.hgt != "183cm" {
        t.Errorf("extractPassport(<1>).hgt, %s given, \"183cm\" expected", given_result.hgt)
    }
    if given_result.hcl != "#fffffd" {
        t.Errorf("extractPassport(<1>).hcl, %s given, \"#fffffd\" expected", given_result.hcl)
    }
    if given_result.ecl != "gry" {
        t.Errorf("extractPassport(<1>).ecl, %s given, \"gry\" expected", given_result.ecl)
    }
    if given_result.pid != "860033327" {
        t.Errorf("extractPassport(<1>).pid, %s given, \"860033327\" expected", given_result.pid)
    }
    if given_result.cid != "147" {
        t.Errorf("extractPassport(<1>).pid, %s given, \"147\" expected", given_result.cid)
    }
  }
}

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

func TestCheckPolicyTwo(t *testing.T) {
  var givenResult bool
  givenResult = checkPolicyTwo(1, 3, 'a', "abcde")
  if !givenResult {
    t.Error("checkPolicyTwo(1, 3, 'a', \"abcde\") gives false, true expected")
  }
  givenResult = checkPolicyTwo(1, 3, 'b', "cdefg")
  if givenResult {
    t.Error("checkPolicyTwo(1, 3, 'b', \"cdefg\") gives true, false expected")
  }
  givenResult = checkPolicyTwo(2, 9, 'c', "ccccccccc")
  if givenResult {
    t.Error("checkPolicyTwo(1, 3, 'c', \"ccccccccc\") gives true, false expected")
  }
}

func TestConvertLine(t *testing.T) {
  var (
    givenNumA int
    givenNumB int
    givenLetter byte
    givenPassword string
    givenError error
  )
  givenNumA, givenNumB, givenLetter, givenPassword, givenError = convertLine("1-3 a: abcde")
  if givenError != nil {
    t.Errorf("convertLine(\"1-3 a: abcde\") return err %s", givenError)
  } else if givenNumA != 1 || givenNumB != 3 || givenLetter != 'a' || givenPassword != "abcde" {
    t.Errorf("convertLine(\"1-3 a: abcde\") return unexpected values (%v, %v, %c, %s)",
             givenNumA, givenNumB, givenLetter, givenPassword)
  }
  givenNumA, givenNumB, givenLetter, givenPassword, givenError = convertLine("1-3 b: cdefg")
  if givenError != nil {
    t.Errorf("convertLine(\"1-3 b: cdefg\") return err %s", givenError)
  } else if givenNumA != 1 || givenNumB != 3 || givenLetter != 'b' || givenPassword != "cdefg" {
    t.Errorf("convertLine(\"1-3 b: cdefg\") return unexpected values (%v, %v, %c, %s)",
             givenNumA, givenNumB, givenLetter, givenPassword)
  }
  givenNumA, givenNumB, givenLetter, givenPassword, givenError = convertLine("2-9 c: ccccccccc")
  if givenError != nil {
    t.Errorf("convertLine(\"2-9 a: ccccccccc\") return err %s", givenError)
  } else if givenNumA != 2 || givenNumB != 9 || givenLetter != 'c' || givenPassword != "ccccccccc" {
    t.Errorf("convertLine(\"2-9 a: ccccccccc\") return unexpected values (%v, %v, %c, %s)",
             givenNumA, givenNumB, givenLetter, givenPassword)
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
