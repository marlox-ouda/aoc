#!/usr/bin/env python3
# -*- coding: utf-8 -*-

BANNED_STRINGS = ('ab', 'cd', 'pq', 'xy')


def validator_voyels(evaluated_string):
    return sum(evaluated_string.count(voyel) for voyel in 'aeiou') > 2


def validator_double(evaluated_string):
    return any(first_letter == second_letter for first_letter, second_letter in
               zip(evaluated_string[:-1], evaluated_string[1:]))


def validator_exclusion(evaluated_string):
    return not any(banned_string in evaluated_string
                   for banned_string in BANNED_STRINGS)


def validator_repeated_pair(evaluated_string):
    return any(
        evaluated_string.count(first_letter + second_letter) > 1
        for first_letter, second_letter in zip(evaluated_string[:-1],
                                               evaluated_string[1:]))


def validator_sandwitch_repeat(evaluated_string):
    return any(first_letter == third_letter
               for first_letter, _, third_letter in zip(
                   evaluated_string[:-2], evaluated_string[1:-1],
                   evaluated_string[2:]))


OLD_VALIDATORS = (validator_voyels, validator_double, validator_exclusion)
NEW_VALIDATORS = (validator_repeated_pair, validator_sandwitch_repeat)


def validator_all(evaluated_string):
    return all(validator(evaluated_string) for validator in NEW_VALIDATORS)


def main_manuel():
    string_to_eval = input("Saisie manuelle : ")
    print(validator_all(string_to_eval))


def main_auto():
    with open('input.txt') as fd:
        strings_to_eval = fd.readlines()
    strings_to_eval = filter(validator_all, strings_to_eval)
    print(len(list(strings_to_eval)))


if __name__ == '__main__':
    main_auto()
