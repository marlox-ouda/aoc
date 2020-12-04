#!/usr/bin/env python3
# -*- coding: utf-8 -*-


def char_to_num(converted_char):
    return ord(converted_char) - ord('a')


def str_to_num_list(converted_str):
    num_list = list()
    for char in converted_str:
        num_list.append(char_to_num(char))
    return num_list


def num_list_to_str(num_list):
    num_list = (char + ord('a') for char in num_list)
    num_list = (chr(char) for char in num_list)
    return ''.join(num_list)


def iter_password(password_elements):
    """Iter through potential password"""
    while True:
        for pos in range(7, 0, -1):
            if password_elements[pos] == 25:
                password_elements[pos] = 0
            else:
                password_elements[pos] += 1
                break
        yield password_elements


def validate_double(pwd):
    tested_passwords = (first == second and first or False
                        for first, second in zip(pwd[:-1], pwd[1:]))
    valide_letters = list(filter(lambda l: l, tested_passwords))
    if len(list(valide_letters)) < 2:
        return False
    if valide_letters[0] == valide_letters[-1]:
        return False
    return True


VALIDATORS = (
    lambda pwd: not char_to_num('i') in pwd,
    lambda pwd: not char_to_num('o') in pwd,
    lambda pwd: not char_to_num('l') in pwd,
    lambda pwd: any(first + 1 == second and second + 1 == third for first,
                    second, third in zip(pwd[:-2], pwd[1:-1], pwd[2:])),
    validate_double,
)


def validate_password(tryied_password):
    """Check if password is correct"""
    if all(validator(tryied_password) for validator in VALIDATORS):
        return True
    return False


def next_password(current_password):
    """Define the new password"""
    for tryied_password in iter_password(current_password):
        if validate_password(tryied_password):
            break
    return tryied_password


def main():
    """Main function"""
    current_password = str_to_num_list('vzbxkghb')
    new_password = next_password(current_password)
    print(num_list_to_str(new_password))
    new_password = next_password(new_password)
    print(num_list_to_str(new_password))


if __name__ == '__main__':
    main()
