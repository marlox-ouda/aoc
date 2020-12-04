#!/usr/bin/env python3
# -*- coding: utf-8 -*-
import re

HEXA_RE = re.compile(r'\\x[0-9a-fA-F]{2}')


class CharacterParser:
    def __init__(self, characters):
        self.characters = characters

    def __len__(self):
        return sum((
            self.characters.startswith('"') and 1 or 0,
            self.characters.endswith('"') and 1 or 0,
            self.characters.count("\\"),
            #self.characters.count("\\\\"),
            #self.characters.count("\\\""),
            2 * len(HEXA_RE.findall(self.characters)),
        ))


def main():
    with open('input.txt') as fd:
        instructions = fd.readlines()
    instructions = (CharacterParser(instruction[:-1])
                    for instruction in instructions)
    instructions = (len(parser) for parser in instructions)
    print(sum(instructions))


if __name__ == '__main__':
    main()
