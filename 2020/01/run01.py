#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import functools
import itertools
import operator
import pathlib

def run(group_by: int=2):
    input_file = pathlib.Path('input1.txt')
    input_file_data = input_file.read_text()
    input_str_lines = input_file_data.split('\n')
    input_str_lines = filter(operator.truth, input_str_lines)
    input_integers = map(int, input_str_lines)
    integers_combinaison = itertools.combinations(input_integers, group_by)
    integers_combinaison = filter(lambda combinaison: sum(combinaison) == 2020,
                                  integers_combinaison)
    valid_combinaison = next(integers_combinaison)
    expense_fix = functools.reduce(operator.mul, valid_combinaison, 1)
    print(expense_fix)

run()
run(3)
