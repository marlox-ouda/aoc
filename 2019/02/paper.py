#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Box:
    def __init__(self, length, height, width):
        self.length = length
        self.height = height
        self.width = width

    @property
    def paper_area(self):
        sides_area = (
            self.length * self.height,
            self.length * self.width,
            self.height * self.width,
        )
        return 2 * sum(sides_area) + min(sides_area)

    @property
    def volume(self):
        return self.length * self.height * self.width

    @property
    def smallest_perimeter(self):
        ordered_dimensions = sorted((
            self.length,
            self.height,
            self.width,
        ))
        ordered_dimensions = list(ordered_dimensions)
        return 2 * (ordered_dimensions[0] + ordered_dimensions[1])

    @property
    def ribon_size(self):
        return self.smallest_perimeter + self.volume

    @classmethod
    def from_str(cls, dimensions):
        length, height, width = map(int, dimensions.split('x'))
        return Box(length, height, width)


def main_manuel():
    dimensions = input("Saisie manuelle : ")
    box = Box.from_str(dimensions)
    print(box.paper_area)


def main_auto():
    with open('input.txt') as fd:
        liste_dimensions = fd.readlines()
    area = 0
    ribon = 0
    for dimensions in liste_dimensions:
        box = Box.from_str(dimensions)
        area += box.paper_area
        ribon += box.ribon_size
    print(area)
    print(ribon)


if __name__ == '__main__':
    main_auto()
