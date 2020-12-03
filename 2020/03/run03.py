#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def count_trees(right_step:'int', down_step:'int'):
    print(f"{right_step}, {down_step}")
    input_file = pathlib.Path('input3.txt')
    input_file_data = input_file.read_text()
    input_str_lines = input_file_data.split('\n')
    input_str_lines = filter(operator.truth, input_str_lines)
    input_str_lines = list(input_str_lines)
    input_lines = enumerate(input_str_lines)
    input_lines = filter(lambda enum_line: enum_line[0] % down_step == 0, input_lines)
    input_lines = (line for pos, line in input_lines)
    input_lines = enumerate(input_lines)
    input_lines = (line[(right_step*pos)%len(line)] for pos, line in input_lines)
    input_lines = list(input_lines)
    input_lines = filter(lambda obj: obj == '#', input_lines)
    number_trees = len(list(input_lines))
    print(number_trees)
    return number_trees

def run_one():
    print(count_trees(3, 1))

def run_two():
    combinaisons =(
        (1, 1),
        (3, 1),
        (5, 1),
        (7, 1),
        (1, 2),
    )
    number_trees = (count_trees(*combo) for combo in combinaisons)
    multiply = functools.reduce(operator.mul, number_trees, 1)
    print(multiply)

run_one()
run_two()
