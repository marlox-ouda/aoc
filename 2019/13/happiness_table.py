#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from collections import defaultdict
from itertools import permutations
from re import compile

FEELING_RE = compile(
    r'^(?P<person>\w+) would (?P<dir>lose|gain) (?P<value>\d+) '
    r'happiness units by sitting next to (?P<other>\w+).$')


class HappinessDB:
    def __init__(self):
        self.persons = set()
        self.happinesses = defaultdict(lambda: dict())

    def add_feeling(self, person, other, happiness_variation):
        self.persons.add(person)
        self.persons.add(other)
        self.happinesses[person][other] = happiness_variation

    def better_placement(self):
        better_happiness = None
        for proposed_order in permutations(self.persons, len(self.persons)):
            happiness = 0
            for left, right in zip(proposed_order[:-1], proposed_order[1:]):
                happiness += self.happinesses[left][right]
                happiness += self.happinesses[right][left]
            happiness += self.happinesses[proposed_order[0]][
                proposed_order[-1]]
            happiness += self.happinesses[proposed_order[-1]][
                proposed_order[0]]
            if better_happiness is None or happiness > better_happiness:
                print(proposed_order)
                better_happiness = happiness
        return better_happiness

    def better_placement_with_you(self):
        better_happiness = None
        for proposed_order in permutations(self.persons, len(self.persons)):
            happiness = 0
            for left, right in zip(proposed_order[:-1], proposed_order[1:]):
                happiness += self.happinesses[left][right]
                happiness += self.happinesses[right][left]
            if better_happiness is None or happiness > better_happiness:
                print(proposed_order)
                better_happiness = happiness
        return better_happiness


def main():
    happiness_db = HappinessDB()
    with open('input.txt') as fd:
        feelings = fd.readlines()
    feelings = (feeling.rstrip('\n') for feeling in feelings)
    feelings = (FEELING_RE.match(feeling) for feeling in feelings)
    feelings = filter(lambda x: x, feelings)
    for feeling in feelings:
        happiness_variation = int(feeling.group('value'))
        if feeling.group('dir') == 'lose':
            happiness_variation = -happiness_variation
        happiness_db.add_feeling(feeling.group('person'),
                                 feeling.group('other'), happiness_variation)
    print(happiness_db.better_placement())
    print(happiness_db.better_placement_with_you())


if __name__ == '__main__':
    main()
