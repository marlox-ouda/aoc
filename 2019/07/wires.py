#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import re

OP_DIRECT_RE = re.compile(r'^(?P<in>\w+) -> (?P<out>\w+)$')
OP_NOT_RE = re.compile(r'^NOT (?P<in>\w+) -> (?P<out>\w+)$')
OP_AND_RE = re.compile(r'^(?P<in1>\w+) AND (?P<in2>\w+) -> (?P<out>\w+)$')
OP_OR_RE = re.compile(r'^(?P<in1>\w+) OR (?P<in2>\w+) -> (?P<out>\w+)$')
OP_LSHIFT_RE = re.compile(
    r'^(?P<in1>\w+) LSHIFT (?P<in2>\w+) -> (?P<out>\w+)$')
OP_RSHIFT_RE = re.compile(
    r'^(?P<in1>\w+) RSHIFT (?P<in2>\w+) -> (?P<out>\w+)$')

OP_REGEX = (
    (OP_DIRECT_RE, 'op_direct'),
    (OP_NOT_RE, 'op_not'),
    (OP_AND_RE, 'op_and'),
    (OP_OR_RE, 'op_or'),
    (OP_LSHIFT_RE, 'op_lshift'),
    (OP_RSHIFT_RE, 'op_rshift'),
)


class Wire:
    def __init__(self):
        self.value = None
        self.formula = None
        self.dependant_wires = list()
        self.nb_dependancies = 0

    def __str__(self):
        return str(self.value)


class Board:
    def __init__(self):
        self.wires = dict()

    def set_wire_value(self, wire_id, value):
        if wire_id not in self.wires:
            self.wires[wire_id] = Wire()
        wire = self.wires[wire_id]
        if wire.value is not None:
            return
        wire.value = value
        wire.nb_dependancies = 0
        for dependant_wire_uid in wire.dependant_wires:
            dependant_wire = self.wires[dependant_wire_uid]
            dependant_wire.nb_dependancies -= 1
            if not dependant_wire.nb_dependancies:
                new_value = dependant_wire.formula(self)
                self.set_wire_value(dependant_wire_uid, new_value)

    def ensure_out(self, output_uid):
        if output_uid not in self.wires:
            self.wires[output_uid] = Wire()

    def make_dependancy(self, provider_wire_uid, dependant_wire_uid):
        if provider_wire_uid not in self.wires:
            self.wires[provider_wire_uid] = Wire()
        if self.wires[provider_wire_uid].value is None:
            self.wires[dependant_wire_uid].nb_dependancies += 1
            self.wires[provider_wire_uid].dependant_wires.append(
                dependant_wire_uid)

    def get_lazy_input(self, input_value, output_uid):
        if input_value.isdigit():
            lazy_value = lambda board: int(input_value)
        else:
            self.make_dependancy(input_value, output_uid)
            lazy_value = lambda board: self.wires[input_value].value
        return lazy_value

    def op_direct(self, in1, out):
        self.ensure_out(out)
        in_value = self.get_lazy_input(in1, out)
        self.wires[out].formula = lambda board: in_value(board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def op_not(self, in1, out):
        self.ensure_out(out)
        in_value = self.get_lazy_input(in1, out)
        self.wires[out].formula = lambda board: ~in_value(board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def op_and(self, in1, in2, out):
        self.ensure_out(out)
        in1_value = self.get_lazy_input(in1, out)
        in2_value = self.get_lazy_input(in2, out)
        self.wires[out].formula = lambda board: in1_value(board) & in2_value(
            board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def op_or(self, in1, in2, out):
        self.ensure_out(out)
        in1_value = self.get_lazy_input(in1, out)
        in2_value = self.get_lazy_input(in2, out)
        self.wires[out].formula = lambda board: in1_value(board) | in2_value(
            board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def op_lshift(self, in1, in2, out):
        self.ensure_out(out)
        in1_value = self.get_lazy_input(in1, out)
        in2_value = self.get_lazy_input(in2, out)
        self.wires[out].formula = lambda board: in1_value(board) << in2_value(
            board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def op_rshift(self, in1, in2, out):
        self.ensure_out(out)
        in1_value = self.get_lazy_input(in1, out)
        in2_value = self.get_lazy_input(in2, out)
        self.wires[out].formula = lambda board: in1_value(board) >> in2_value(
            board)
        if not self.wires[out].nb_dependancies:
            self.set_wire_value(out, self.wires[out].formula(self))

    def parse_instruction(self, instruction):
        for regex, method_name in OP_REGEX:
            match = regex.fullmatch(instruction)
            if match:
                method = getattr(self, method_name)
                method(*match.groups())
                break
        else:
            print(instruction)


def main():
    board = Board()
    board.wires['b'] = Wire()
    board.wires['b'].value = 3176
    with open('input.txt') as fd:
        instructions = fd.readlines()
    for instruction in instructions:
        board.parse_instruction(instruction[:-1])
    print(board.wires['a'])


if __name__ == '__main__':
    main()
