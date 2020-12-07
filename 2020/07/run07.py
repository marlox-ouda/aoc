#!/usr/bin/env python3
# -*- coding: utf-8 -*-

import collections
import functools
import itertools
import operator
import pathlib

from pprint import pprint as print

def get_content(bag_content):
    if bag_content == 'no other bags':
        return None
    bag_content = bag_content.split(', ')
    bag_content = (content.rstrip('s')[:-4] for content in bag_content)
    bag_content = (content.split(' ', 1) for content in bag_content)
    bag_content = ((int(number), subbag) for number, subbag in bag_content)
    return tuple(bag_content)

def get_bags():
    input_file = pathlib.Path('input7.txt')
    input_file_data = input_file.read_text()
    bag_lines = input_file_data.split('\n')
    bag_lines = filter(operator.truth, bag_lines)
    bag_lines = (bag_line.rstrip('.') for bag_line in bag_lines)
    bag_lines = (bag_line.split(' bags contain ', 1) for bag_line in bag_lines)
    bag_dict = {bag: get_content(content) for bag, content in bag_lines}
    return bag_dict

def run_one():
    bag_dict = get_bags()
    bag_dict = {bag: content is not None and tuple((content_bag[1] for content_bag in content)) or None for bag, content in bag_dict.items()}
    bags = collections.deque(bag_dict.keys())
    while bags:
        bag = bags.popleft()
        content = bag_dict[bag]
        if content is None:
            pass
        elif 'shiny gold' in content:
            bag_dict[bag] = True
        else:
            if any((bag_dict[content_bag] == True for content_bag in content)):
                bag_dict[bag] = True
            else:
                content = tuple((content_bag for content_bag in content if bag_dict[content_bag] is not None))
                if content:
                    bag_dict[bag] = content
                    bags.append(bag)
                else:
                    bag_dict[bag] = None
    bags = (bag for bag, content in bag_dict.items() if content == True)
    print(len(list(bags)))

def run_two():
    bag_dict = get_bags()
    bags = collections.deque([(1, 'shiny gold')])
    total = 0
    while bags:
        number, bag = bags.pop()
        print(f"{number}×{bag}")
        total += number
        content_bag = bag_dict[bag]
        if content_bag is not None:
            for sub_bag_number, sub_bag in content_bag:
                bags.append((number * sub_bag_number, sub_bag))
    print(total)



run_one()
run_two()
