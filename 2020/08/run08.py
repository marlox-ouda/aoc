#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import collections
import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_instructions():
    input_file = pathlib.Path('input8.txt')
    input_file_data = input_file.read_text()
    instructions = input_file_data.split('\n')
    instructions = filter(operator.truth, instructions)
    instructions = (instr.split(' ', 1) for instr in instructions)
    instructions = ((instr, int(value))  for instr, value in instructions)
    return instructions

def run(instructions):
    accumulator = 0
    position = 0
    positions = set()
    final = False
    while True:
        instr, value = instructions[position]
        if instr == 'jmp':
            position += value
        else:
            position += 1
            if instr == 'acc':
                accumulator += value
        if position in positions:
            break
        if position == len(instructions):
            final = True
            break
        if not (0 <= position < len(instructions)):
            break
        positions.add(position)
    return accumulator, final

def run_one():
    accumulator, _ = run(list(get_instructions()))
    print(accumulator)

def run_two():
    instructions = list(get_instructions())
    for idx, instr_value in enumerate(instructions):
        instr, value = instr_value
        if instr == 'acc':
            continue
        instructions[idx] = (instr == 'nop' and 'jmp' or 'nop', value)
        accumulator, final = run(instructions)
        instructions[idx] = instr_value
        if final:
            print(accumulator)
            break
    accumulator, _ = run(list(get_instructions()))

    pass

run_one()
run_two()
