#!/usr/bin/env python3
# -*- coding: utf-8 -*-


class Person:
    def __init__(self, hit_points):
        self.hit_points = hit_points

    @property
    def is_death(self):
        return self.hit_points <= 0

    def attack(self, other_person):
        damage = max(1, self.damage - other_person.armor)
        other_person.hit_points -= damage

    def win_fight_against(self, other_person):
        while True:
            self.attack(other_person)
            if other_person.is_death:
                return True
            other_person.attack(self)
            if self.is_death:
                return False


class PJ(Person):
    def __init__(self):
        super().__init__(100)
        self.equipments = list()

    @property
    def damage(self):
        return sum(equipment.damage_mod for equipment in self.equipments)

    @property
    def armor(self):
        return sum(equipment.armor_mod for equipment in self.equipments)

    @property
    def spend_money(self):
        return sum(equipment.cost for equipment in self.equipments)


class Boss(Person):
    def __init__(self, hit_points, damage, armor):
        super().__init__(hit_points)
        self.damage = damage
        self.armor = armor


class Equipment:
    def __init__(self, category, name, cost, damage_mod, armor_mod):
        self.category = category
        self.name = name
        self.cost = cost
        self.damage_mod = damage_mod
        self.armor_mod = armor_mod

    def __repr__(self):
        return self.name


EQUIPEMENTS = (
    Equipment('weapon', 'dagger', 8, 4, 0),
    Equipment('weapon', 'shortsword', 10, 5, 0),
    Equipment('weapon', 'warhammer', 25, 6, 0),
    Equipment('weapon', 'longsword', 40, 7, 0),
    Equipment('weapon', 'greataxe', 74, 8, 0),
    Equipment('armor', 'leather', 13, 0, 1),
    Equipment('armor', 'chainmail', 31, 0, 2),
    Equipment('armor', 'splintmail', 53, 0, 3),
    Equipment('armor', 'bandedmail', 75, 0, 4),
    Equipment('armor', 'platemail', 102, 0, 5),
    Equipment('ring', 'damage+1', 25, 1, 0),
    Equipment('ring', 'damage+2', 50, 2, 0),
    Equipment('ring', 'damage+3', 100, 3, 0),
    Equipment('ring', 'armor+1', 20, 0, 1),
    Equipment('ring', 'armor+2', 40, 0, 2),
    Equipment('ring', 'armor+3', 80, 0, 3),
)


def generate_pj():
    weapons = [
        equipment for equipment in EQUIPEMENTS
        if equipment.category == 'weapon'
    ]
    armors = [
        equipment for equipment in EQUIPEMENTS if equipment.category == 'armor'
    ]
    armors.append(None)
    rings = [
        equipment for equipment in EQUIPEMENTS if equipment.category == 'ring'
    ]
    rings.append(None)
    for weapon in weapons:
        for armor in armors:
            for ring1 in rings:
                for ring2 in rings:
                    if ring1 is not None and ring1 == ring2:
                        continue
                    if ring1 is None and ring2 is not None:
                        continue
                    pj = PJ()
                    pj.equipments.append(weapon)
                    if armor:
                        pj.equipments.append(armor)
                    if ring1:
                        pj.equipments.append(ring1)
                    if ring2:
                        pj.equipments.append(ring2)
                    yield pj


def main():
    pjs = (pj for pj in generate_pj() if pj.win_fight_against(Boss(109, 8, 2)))
    best_pj = min(pjs, key=lambda pj: pj.spend_money)
    print(best_pj.spend_money)
    print(best_pj.equipments)
    print(best_pj.damage)
    print(best_pj.armor)

    pjs = (pj for pj in generate_pj()
           if not pj.win_fight_against(Boss(109, 8, 2)))
    worst_pj = max(pjs, key=lambda pj: pj.spend_money)
    print(worst_pj.spend_money)
    print(worst_pj.equipments)
    print(worst_pj.damage)
    print(worst_pj.armor)


if __name__ == '__main__':
    main()
