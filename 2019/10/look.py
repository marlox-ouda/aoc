#!/usr/bin/env python3
# -*- coding: utf-8 -*-


def look_and_say(number):
    new_number = list()
    prev_digit = None
    nb_occurences = 0
    for cur_digit in number:
        if prev_digit == cur_digit:
            nb_occurences += 1
        else:
            if prev_digit is not None:
                new_number.append(str(nb_occurences))
                new_number.append(prev_digit)
            prev_digit = cur_digit
            nb_occurences = 1
    new_number.append(str(nb_occurences))
    new_number.append(cur_digit)
    return ''.join(new_number)


def test(number):
    new_number = look_and_say(number)
    print(f'{new_number} = look_and_say({number})')


def main():
    test('1')
    test('11')
    test('21')
    test('1211')
    test('111221')
    number = '1113222113'
    for _ in range(40):
        number = look_and_say(number)
    print(len(number))
    for _ in range(10):
        number = look_and_say(number)
    print(len(number))


if __name__ == '__main__':
    main()
