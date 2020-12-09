#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import collections
import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_numbers():
    input_file = pathlib.Path('input9.txt')
    input_file_data = input_file.read_text()
    instructions = input_file_data.split('\n')
    instructions = filter(operator.truth, instructions)
    numbers = map(int, instructions)
    return list(numbers)

def run_one(numbers):
    numbers = iter(numbers)
    current_numbers = collections.deque((next(numbers) for _ in range(25)))
    while True:
        next_number = next(numbers)
        if next_number in map(sum, itertools.combinations(current_numbers, 2)):
            current_numbers.popleft()
            current_numbers.append(next_number)
            continue
        break
    print(next_number)
    return next_number

def run_two(numbers, one_solution):
    numbers = iter(numbers)
    current_numbers = collections.deque()
    total = 0
    while True:
        while total < one_solution:
            next_number = next(numbers)
            total += next_number
            current_numbers.append(next_number)
        if total == one_solution:
            return min(current_numbers) + max(current_numbers)
        while total > one_solution:
            prev_number = current_numbers.popleft()
            total -= prev_number
        if total == one_solution:
            return min(current_numbers) + max(current_numbers)

numbers = get_numbers()
solution = run_one(numbers)
print(run_two(numbers, solution))
