package main

import "testing"

func TestFindPair(t *testing.T) {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  pair_result_a, pair_result_b, err := find_pair(input_numbers)
  if err != nil {
    t.Error("find_pair: Do not find any valid pair on example test")
  } else if (pair_result_a == 1721 && pair_result_b == 299) {
    t.Log("find_pair: Sucess on example test")
  } else if (pair_result_b == 1721 && pair_result_a == 299) {
    t.Log("find_pair: Sucess on example test")
  } else {
    t.Errorf("find_pair: Do not return 1721 and 299 on example test but %v and %v", pair_result_a, pair_result_b)
  }
}

func TestFixExpense(t *testing.T) {
  input_numbers := make([]int, 6)
  input_numbers[0] = 1721
  input_numbers[1] = 979
  input_numbers[2] = 366
  input_numbers[3] = 299
  input_numbers[4] = 675
  input_numbers[5] = 1456
  given_multiple, err := fix_expense(input_numbers)
  expected_multiple := 514579
  if err != nil {
    t.Error("fix_expense: Do not find any valid pair on example test")
  } else if (given_multiple == expected_multiple) {
    t.Log("fix_expense: Sucess on example test")
  } else {
    t.Errorf("fix_expense: Do not return %v on example test but %v",
             expected_multiple, given_multiple)
  }
}
