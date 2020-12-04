#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Elevator:
    def __init__(self, starting_level=0):
        self._level = starting_level
        self.move_counter = 0
        self._first_in_basement = None

    @property
    def level(self):
        return self._level

    @property
    def first_in_basement(self):
        return self._first_in_basement

    def read_instruction(self, instruction):
        if instruction == '(':
            self._level += 1
        elif instruction == ')':
            self._level -= 1
        else:
            print("'{}'' is not a valid instruction for "
                  "the elevator", instruction)
        self.move_counter += 1
        if self._first_in_basement is None and self._level < 0:
            self._first_in_basement = self.move_counter

    @classmethod
    def read_instructions(cls, instructions):
        elevator = cls()
        for instruction in instructions:
            elevator.read_instruction(instruction)
        return elevator.first_in_basement


def main_manuel():
    instructions = input("Saisie manuelle : ")
    print(Elevator.read_instructions(instructions))


def main_auto():
    with open('input.txt') as fd:
        instructions = fd.readline()
    print(Elevator.read_instructions(instructions))


if __name__ == '__main__':
    main_auto()
