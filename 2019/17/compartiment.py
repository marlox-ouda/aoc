#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from re import compile


class Container:
    def __init__(self, uid, capacity):
        self.uid = 1 << uid
        self.capacity = capacity

    #def __repr__(self):
    #    return f"{self.uid}:{self.capacity}L"


class Arrangement:
    def __init__(self, containers, repartition):
        self.containers = list()
        for container in containers:
            if container.uid & repartition:
                self.containers.append(container)

    @property
    def capacity(self):
        return sum(container.capacity for container in self.containers)


def repartition_generator(containers):
    nb_containers = len(containers)
    yield from range(pow(2, nb_containers))


def main():
    with open('input.txt') as fd:
        containers = fd.readlines()
    containers = (ing.rstrip('\n') for ing in containers)
    containers = filter(lambda x: x, containers)
    containers = map(int, containers)
    containers = map(lambda x: Container(x[0], x[1]), enumerate(containers))
    containers = list(containers)

    arrangements = (Arrangement(containers, repartition)
                    for repartition in repartition_generator(containers))
    arrangements = (1 for arrangement in arrangements
                    if arrangement.capacity == 150)
    print(sum(arrangements))


if __name__ == '__main__':
    main()
