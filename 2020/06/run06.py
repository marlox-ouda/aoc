#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_declarations():
    input_file = pathlib.Path('input6.txt')
    input_file_data = input_file.read_text()
    group_declarations = input_file_data.split('\n\n')
    return group_declarations

def run_one():
    group_declarations = get_declarations()
    group_declarations = (declaration.replace('\n', '') for declaration in group_declarations)
    group_declarations = map(set, group_declarations)
    group_declarations = map(len, group_declarations)
    print(sum(group_declarations))

def run_two():
    group_declarations = get_declarations()
    group_declarations = (declaration.strip('\n') for declaration in group_declarations)
    group_declarations = (map(set, declaration.split('\n')) for declaration in group_declarations)
    group_declarations = (set.intersection(*declaration) for declaration in group_declarations)
    group_declarations = map(len, group_declarations)
    print(sum(group_declarations))

run_one()
run_two()
