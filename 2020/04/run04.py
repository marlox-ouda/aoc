#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import functools
import itertools
import operator
import pathlib
import re

from pprint import pprint as print

COLOR_RE = re.compile(r'#[0-9a-f]{6}')

def extract_password_data(password_line: str) -> 'dict[str,str]':
    password_items = password_line.split(' ')
    password_items = (item.split(':', 1) for item in password_items)
    password_items = {k: v for k, v in password_items}
    return password_items

def read_passwords():
    input_file = pathlib.Path('input4.txt')
    input_file_data = input_file.read_text()
    password_blocks = input_file_data.split('\n\n')
    password_blocks = (line.strip('\n') for line in password_blocks)
    password_lines = (line.replace('\n', ' ') for line in password_blocks)
    password_items = map(extract_password_data, password_lines)
    return password_items

REQUIRED_FIELDS = {
    'byr': lambda v: 1920 <= int(v) <= 2002,
    'iyr': lambda v: 2010 <= int(v) <= 2020,
    'eyr': lambda v: 2020 <= int(v) <= 2030,
    'hgt': lambda v: (v.endswith('cm') and 150 <= int(v[:-2]) <= 193) or (v.endswith('in') and 59 <= int(v[:-2]) <= 76),
    'hcl': lambda v: COLOR_RE.fullmatch(v),
    'ecl': lambda v: v in ('amb', 'blu', 'brn', 'gry', 'grn', 'hzl', 'oth'),
    'pid': lambda v: len(v) == 9 and v.isdigit(),
}

def one():
    passwords = read_passwords()
    passwords = filter(lambda password: all((field in password for field in REQUIRED_FIELDS)), passwords)
    print(len(list(passwords)))

def two():
    passwords = read_passwords()
    passwords = filter(lambda password: all((field in password for field in REQUIRED_FIELDS)), passwords)
    passwords = filter(lambda password: all((func(password[field]) for field, func in REQUIRED_FIELDS.items())), passwords)
    print(len(list(passwords)))

one()
two()
