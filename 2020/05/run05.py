#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_seat_numbers():
    input_file = pathlib.Path('input5.txt')
    input_file_data = input_file.read_text()
    input_str_lines = input_file_data.split('\n')
    input_str_lines = filter(operator.truth, input_str_lines)
    bin_lines = ((symbol in ('B', 'R') and '1' or '0'
                  for symbol in line
                  )
                 for line in input_str_lines)
    bin_lines = (''.join(line_symbols) for line_symbols in bin_lines)
    seat_numbers = (int(bin_entry, 2) for bin_entry in bin_lines)
    return seat_numbers

def run_one():
    print(max(get_seat_numbers()))

def run_two():
    seats = sorted(get_seat_numbers())
    diff = zip(seats[:-1], seats[1:])
    diff = (seat + 1 for seat, next_seat in diff if next_seat - seat == 2)
    print(list(diff))

run_one()
run_two()
