package main

import "testing"

func TestCheckPolicyOne(t *testing.T) {
  var givenResult bool
  givenResult = checkPolicyOne(1, 3, 'a', "abcde")
  if !givenResult {
    t.Error("checkPolicyOne(1, 3, 'a', \"abcde\") give false, true expected")
  }
  givenResult = checkPolicyOne(1, 3, 'b', "cdefg")
  if givenResult {
    t.Error("checkPolicyOne(1, 3, 'b', \"cdefg\") gives true, false expected")
  }
  givenResult = checkPolicyOne(2, 9, 'c', "ccccccccc")
  if !givenResult {
    t.Error("checkPolicyOne(1, 3, 'c', \"ccccccccc\") gives false, true expected")
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

func TestFindPair(t *testing.T) {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  pair_result_a, pair_result_b, err := findPair(input_numbers)
  if err != nil {
    t.Error("findPair: Do not find any valid pair on example test")
  } else if (pair_result_a == 1721 && pair_result_b == 299) {
    t.Log("findPair: Sucess on example test")
  } else if (pair_result_b == 1721 && pair_result_a == 299) {
    t.Log("findPair: Sucess on example test")
  } else {
    t.Errorf("findPair: Do not return 1721 and 299 on example test but %v and %v", pair_result_a, pair_result_b)
  }
}

func TestFixExpense(t *testing.T) {
  given_multiple := fixExpense(1721, 299)
  expected_multiple := 514579
  if (given_multiple == expected_multiple) {
    t.Log("fixExpense: Sucess on example one test")
  } else {
    t.Errorf("fixExpense: Do not return %v on example one test but %v",
             expected_multiple, given_multiple)
  }
  given_multiple = fixExpense(979, 366, 675)
  expected_multiple = 241861950
  if (given_multiple == expected_multiple) {
    t.Log("fixExpense: Sucess on example two test")
  } else {
    t.Errorf("fixExpense: Do not return %v on example two test but %v",
             expected_multiple, given_multiple)
  }
}

func TestReadInput(t *testing.T) {
  input_numbers, err := readInput()
  if err != nil {
    t.Errorf("readInput: fail with %s error", err)
  }
  expected_first_number := 1918
  expected_last_number := 1407
  expected_len := 200
  if input_numbers[0] != expected_first_number {
    t.Errorf("readInput: Not expected first number, %v expected, %v given",
             expected_first_number, input_numbers[0])
  }
  given_len := len(input_numbers)
  if given_len != expected_len {
    t.Errorf("readInput: Not expected slice len, %v expected, %v given",
             expected_len, given_len)
  }
  if input_numbers[given_len-1] != expected_last_number {
    t.Errorf("readInput: Not expected last number, %v expected, %v given",
             expected_last_number, input_numbers[given_len-1])
  }
}

func TestFindThree(t *testing.T) {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  three_a, three_b, three_c, err := findThree(input_numbers)
  if err != nil {
    t.Error("findThree: Do not find any valid pair on example test")
  } else {
    if !(three_a == 979 || three_b == 979 || three_c == 979) {
      t.Errorf("findThree: Do not contain 979 among [%v,%v,%v]", three_a, three_b, three_c)
    }
    if !(three_a == 366 || three_b == 366 || three_c == 366) {
      t.Errorf("findThree: Do not contain 366 among [%v,%v,%v]", three_a, three_b, three_c)
    }
    if !(three_a == 675 || three_b == 675 || three_c == 675) {
      t.Errorf("findThree: Do not contain 675 among [%v,%v,%v]", three_a, three_b, three_c)
    }
  }
}
