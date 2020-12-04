#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from collections import defaultdict
from itertools import permutations
from re import compile

DISTANCE_RE = compile(r'^(?P<src>\w+) to (?P<dst>\w+) = (?P<distance>\d+)$')


class DistancesDB:
    def __init__(self):
        self.cities = set()
        self.distances = defaultdict(lambda: dict())

    def add_distance(self, src_city, dst_city, distance):
        self.cities.add(src_city)
        self.cities.add(dst_city)
        if src_city < dst_city:
            self.distances[src_city][dst_city] = distance
        else:
            self.distances[dst_city][src_city] = distance

    def get_distance(self, src_city, dst_city):
        if src_city < dst_city:
            return self.distances[src_city][dst_city]
        else:
            return self.distances[dst_city][src_city]

    def shortest_distance(self):
        shortest_distance = None
        for proposed_order in permutations(self.cities, len(self.cities)):
            distance = 0
            for src, dst in zip(proposed_order[:-1], proposed_order[1:]):
                distance += self.get_distance(src, dst)
            if shortest_distance is None or distance < shortest_distance:
                print(proposed_order)
                shortest_distance = distance
        return shortest_distance

    def longest_distance(self):
        longest_distance = 0
        for proposed_order in permutations(self.cities, len(self.cities)):
            distance = 0
            for src, dst in zip(proposed_order[:-1], proposed_order[1:]):
                distance += self.get_distance(src, dst)
            if distance > longest_distance:
                print(proposed_order)
                longest_distance = distance
        return longest_distance


def main():
    distance_db = DistancesDB()
    with open('input.txt') as fd:
        distances = fd.readlines()
    distances = (distance.rstrip('\n') for distance in distances)
    distances = (DISTANCE_RE.match(distance) for distance in distances)
    distances = filter(lambda x: x, distances)
    for distance in distances:
        distance_db.add_distance(distance.group('src'), distance.group('dst'),
                                 int(distance.group('distance')))
    print(distance_db.shortest_distance())
    print(distance_db.longest_distance())


if __name__ == '__main__':
    main()
