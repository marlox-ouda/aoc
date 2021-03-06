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

func TestReadInput(t *testing.T) {
  input_lines, err := readInput()
  if err != nil {
    t.Errorf("readInput: fail with %s error", err)
  }
  expected_first_line := "7-9 l: vslmtglbc"
  expected_last_line := "9-13 p: bppxpjpmpwcpppdprpp"
  expected_len := 1000
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
