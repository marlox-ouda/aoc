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
    t.Error("FindPair: Do not find any valid pair on example test")
  } else if (pair_result_a == 1721 && pair_result_b == 299) {
    t.Log("FindPair: Sucess on example test")
  } else if (pair_result_b == 1721 && pair_result_a == 299) {
    t.Log("FindPair: Sucess on example test")
  } else {
    t.Errorf("FindPair: Do not return 1721 and 299 on example test but %v and %v", pair_result_a, pair_result_b)
  }
}
