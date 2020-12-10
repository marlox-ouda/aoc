#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import collections
import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_numbers():
    input_file = pathlib.Path('input10.txt')
    input_file_data = input_file.read_text()
    instructions = input_file_data.split('\n')
    instructions = filter(operator.truth, instructions)
    numbers = map(int, instructions)
    return sorted(numbers)

def run_one(numbers):
    target = numbers[-1] + 3
    current = 0
    nb_one = 0
    nb_three = 0
    numbers.append(target)
    numbers = iter(numbers)
    while current != target:
        next_number = next(numbers)
        diff = next_number - current
        if diff == 1:
            nb_one += 1
        elif diff == 3:
            nb_three += 1
        elif diff > 3:
            print(f"warning go from {current} to {next_number}")
        current = next_number
    print(nb_one * nb_three)
    return next_number

COMBO_DICT = {
    0:  1,
    1:  1,
    2:  2,
    3:  4,
    4:  7,
}

def run_two(numbers):
    numbers.insert(0, 0)
    numbers_diff = [b - a for a, b in zip(numbers[:-1], numbers[1:])]
    combo = 0
    multi = 1
    for current_diff in numbers_diff:
        if current_diff == 3:
            if combo not in COMBO_DICT:
                COMBO_DICT[combo] = COMBO_DICT[combo-3] + COMBO_DICT[combo-2] + COMBO_DICT[combo-1] + 1
            multi *= COMBO_DICT[combo]
            combo = 0
        elif current_diff == 1:
            combo += 1
    if combo not in COMBO_DICT:
        COMBO_DICT[combo] = COMBO_DICT[combo-3] + COMBO_DICT[combo-2] + COMBO_DICT[combo-1] + 1
    multi *= COMBO_DICT[combo]
    print(multi)

numbers = get_numbers()
run_one(numbers)
run_two(numbers)
