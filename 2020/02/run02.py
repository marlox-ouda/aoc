#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import functools
import itertools
import operator
import pathlib

def common():
    input_file = pathlib.Path('input2.txt')
    input_file_data = input_file.read_text()
    input_str_lines = input_file_data.split('\n')
    input_str_lines = filter(operator.truth, input_str_lines)
    line_components = (line.split(':', 1) for line in input_str_lines)
    line_components = ((policy.split(' '), password[1:]) for policy, password in line_components)
    line_components = ((l[0][0], l[0][1], l[1]) for l in line_components)
    return line_components

def run_one():
    line_components = common()
    line_components = ((tuple(map(int, occurences.split('-'))), password.count(letter)) for occurences, letter, password in line_components)
    line_components = ((l[0][0], l[0][1], l[1]) for l in line_components)
    line_components = (min_nb <= letter_count <= max_nb for min_nb, max_nb, letter_count in line_components)
    line_components = filter(operator.truth, line_components)
    print(len(list(line_components)))

def run_two():
    line_components = common()
    line_components = ((tuple(map(int, position.split('-'))), letter, password) for position, letter, password in line_components)
    line_components = ((l[0][0] - 1, l[0][1] - 1, l[1], l[2]) for l in line_components)
    line_components = ((password[pos_1] == letter, password[pos_2] == letter) for pos_1, pos_2, letter, password in line_components)
    line_components = (chk_1 ^ chk_2 for chk_1, chk_2 in line_components)
    line_components = filter(operator.truth, line_components)
    print(len(list(line_components)))


run_one()
run_two()
