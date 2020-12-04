#!/usr/bin/env python3
# -*- coding: utf-8 -*-

NB_ROWS = 1000
NB_COLS = 1000


class Position:
    def __init__(self, x_pos, y_pos):
        self.x_pos = x_pos
        self.y_pos = y_pos

    @classmethod
    def from_str(cls, position_string):
        x_pos, y_pos = position_string.split(',')
        return cls(int(x_pos), int(y_pos))


class Board:
    def __init__(self):
        self.lights = [False] * NB_ROWS * NB_ROWS

    def iter_rows(self, start, end):
        for row in range(start, end + 1):
            yield row * NB_COLS

    def iter_cols(self, start, end):
        yield from range(start, end + 1)

    def iter_pos(self, start_pos, end_pos):
        for row_value in self.iter_rows(start_pos.x_pos, end_pos.x_pos):
            for col_value in self.iter_cols(start_pos.y_pos, end_pos.y_pos):
                yield row_value + col_value

    def turn_on(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = True

    def turn_off(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = False

    def toggle(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = not self.lights[value]

    def __len__(self):
        return self.lights.count(True)

    def parse_instruction(self, instruction):
        instruction_elements = instruction.split(' ')
        start_pos = Position.from_str(instruction_elements[-3])
        end_pos = Position.from_str(instruction_elements[-1])
        cmd = instruction_elements[-4]
        if cmd == 'on':
            self.turn_on(start_pos, end_pos)
        elif cmd == 'toggle':
            self.toggle(start_pos, end_pos)
        elif cmd == 'off':
            self.turn_off(start_pos, end_pos)


class BoardV2(Board):
    def __init__(self):
        self.lights = [0] * NB_ROWS * NB_ROWS

    def turn_on(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = self.lights[value] + 1

    def turn_off(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = max(0, self.lights[value] - 1)

    def toggle(self, start_pos, end_pos):
        for value in self.iter_pos(start_pos, end_pos):
            self.lights[value] = self.lights[value] + 2

    def __len__(self):
        return sum(self.lights)


def read_manuel_directions(board):
    instruction = input("Saisie manuelle : ")
    board.parse_instruction(instruction)


def read_auto_directions(board):
    with open('input.txt') as fd:
        instructions = fd.readlines()
    for instruction in instructions:
        board.parse_instruction(instruction)


def main():
    board = BoardV2()
    read_auto_directions(board)
    print(len(board))


if __name__ == '__main__':
    main()
