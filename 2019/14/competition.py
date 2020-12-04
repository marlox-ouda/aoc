#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from re import compile

FLYIER_SPECS_RE = compile(
    r'^(?P<name>\w+) can fly (?P<speed>\d+) km/s for (?P<flying_time>\d+) '
    r'seconds, but then must rest for (?P<rest_time>\d+) seconds.$')


class Flyier:
    def __init__(self, name, speed, flying_time, rest_time):
        self.name = name
        self.speed = speed
        self.flying_time = flying_time
        self.rest_time = rest_time

    def distance_at(self, time):
        is_flying = True
        current_time = 0
        current_distance = 0
        while current_time < time:
            if is_flying:
                current_distance += self.speed * min(self.flying_time,
                                                     time - current_time)
                current_time += self.flying_time
            else:
                current_time += self.rest_time
            is_flying = not is_flying
        print(f"{self.name} at {current_distance}km after {time}s")
        return current_distance


class FlyierV2(Flyier):
    def __init__(self, *args, **kwargs):
        super().__init__(*args, **kwargs)
        self.current_distance = 0
        self.score = 0
        self.currently_flying = True
        self.pending_state_time = self.flying_time

    def time_step(self):
        if self.currently_flying:
            self.current_distance += self.speed
        self.pending_state_time -= 1
        if not self.pending_state_time:
            self.currently_flying = not self.currently_flying
            if self.currently_flying:
                self.pending_state_time = self.flying_time
            else:
                self.pending_state_time = self.rest_time

    def __repr__(self):
        return f"{self.name} with {self.score} points"


def main():
    with open('input.txt') as fd:
        flyiers = fd.readlines()
    flyiers = (flyier.rstrip('\n') for flyier in flyiers)
    flyiers = (FLYIER_SPECS_RE.match(flyier) for flyier in flyiers)
    flyiers = filter(lambda x: x, flyiers)
    flyiers = map(lambda flyier: flyier.group, flyiers)
    flyiers = map(
        lambda flyier: FlyierV2(flyier('name'), int(flyier(
            'speed')), int(flyier('flying_time')), int(flyier('rest_time'))),
        flyiers)
    flyiers = list(flyiers)
    flyiers_part1 = map(lambda flyier: flyier.distance_at(2503), flyiers)
    print(max(flyiers_part1))
    for _ in range(2503):
        for flyier in flyiers:
            flyier.time_step()
        max_distance = max(flyier.current_distance for flyier in flyiers)
        for flyier in flyiers:
            if flyier.current_distance == max_distance:
                flyier.score += 1
    for flyier in flyiers:
        print(flyier)
    flyiers_part2 = map(lambda flyier: flyier.score, flyiers)
    print(max(flyiers_part2))


if __name__ == '__main__':
    main()
