#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Position:
    def __init__(self, x_position=0, y_position=0):
        self.x_position = x_position
        self.y_position = y_position

    def __hash__(self):
        return hash((self.x_position, self.y_position))

    def __neg__(self):
        return Position(-self.x_position, -self.y_position)

    def __eq__(self, other_position):
        if isinstance(other_position, Position):
            return self.x_position == other_position.x_position and \
                self.y_position == other_position.y_position
        return False

    def __add__(self, other_position):
        if other_position in DIRECTION_MAPPING:
            return self + DIRECTION_MAPPING[other_position]
        elif isinstance(other_position, Position):
            return Position(
                self.x_position + other_position.x_position,
                self.y_position + other_position.y_position,
            )
        else:
            print("Warn with ", other_position)
            return self


DIRECTION_MAPPING = {
    '<': Position(-1, 0),
    '>': Position(1, 0),
    '^': Position(0, -1),
    'v': Position(0, 1),
}


class HouseMap:
    def __init__(self):
        self.visited_houses = set()

    def deliver(self, position):
        self.visited_houses.add(position)

    @property
    def nb_visited_houses(self):
        return len(self.visited_houses)


class Santa:
    def __init__(self, house_map):
        self.house_map = house_map
        self.position = Position()
        self.deliver()

    def deliver(self):
        self.house_map.deliver(self.position)

    def move(self, direction):
        self.position += direction
        self.deliver()

    def follow(self, directions):
        for direction in directions:
            self.move(direction)


class SantaWithRobot:
    def __init__(self, house_map):
        self.house_map = house_map
        self.santa = Santa(house_map)
        self.robot_santa = Santa(house_map)

    def follow(self, directions):
        santa_turn = True
        for direction in directions:
            if santa_turn:
                self.santa.move(direction)
            else:
                self.robot_santa.move(direction)
            santa_turn = not santa_turn


def read_manuel_directions(santa):
    directions = input("Saisie manuelle : ")
    santa.follow(directions)


def read_auto_directions(santa):
    with open('input.txt') as fd:
        directions = fd.readline()
    santa.follow(directions)


def main():
    house_map = HouseMap()
    santa = SantaWithRobot(house_map)
    read_auto_directions(santa)
    print(house_map.nb_visited_houses)


if __name__ == '__main__':
    main()
