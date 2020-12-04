#!/usr/bin/env python3
# -*- coding: utf-8 -*-

from re import compile

INGREDIENT_SPECS_RE = compile(
    r'^(?P<name>\w+): capacity (?P<capacity>-?\d+), '
    r'durability (?P<durability>-?\d+), flavor (?P<flavor>-?\d+), '
    r'texture (?P<texture>-?\d+), calories (?P<calories>-?\d+)')


class Ingredient:
    def __init__(self, name, capacity, durability, flavor, texture, calories):
        self.name = name
        self.capacity = capacity
        self.durability = durability
        self.flavor = flavor
        self.texture = texture
        self.calories = calories

    def __repr__(self):
        return self.name

    @classmethod
    def from_spec_re(cls, re_named_group):
        capacity = int(re_named_group('capacity'))
        durability = int(re_named_group('durability'))
        flavor = int(re_named_group('flavor'))
        texture = int(re_named_group('texture'))
        calories = int(re_named_group('calories'))
        return Ingredient(re_named_group('name'), capacity, durability, flavor,
                          texture, calories)


class IngredientUsage:
    def __init__(self, ingredient, quantity):
        self.ingredient = ingredient
        self.quantity = quantity

    @property
    def name(self):
        return self.ingredient.name

    @property
    def capacity(self):
        return self.ingredient.capacity * self.quantity

    @property
    def durability(self):
        return self.ingredient.durability * self.quantity

    @property
    def flavor(self):
        return self.ingredient.flavor * self.quantity

    @property
    def texture(self):
        return self.ingredient.texture * self.quantity

    @property
    def calories(self):
        return self.ingredient.calories * self.quantity

    def __repr__(self):
        return f"{self.quantity}x{self.name}"


class Recipe:
    def __init__(self, ingredients, repartition):
        self.ingredients = list()
        for ingredient in ingredients:
            self.ingredients.append(
                IngredientUsage(ingredient, repartition[ingredient.name]))

    @property
    def capacity(self):
        return max(0, sum(ing.capacity for ing in self.ingredients))

    @property
    def durability(self):
        return max(0, sum(ing.durability for ing in self.ingredients))

    @property
    def flavor(self):
        return max(0, sum(ing.flavor for ing in self.ingredients))

    @property
    def texture(self):
        return max(0, sum(ing.texture for ing in self.ingredients))

    @property
    def calories(self):
        return max(0, sum(ing.calories for ing in self.ingredients))

    @property
    def score(self):
        return self.capacity * self.durability * self.flavor * self.texture

    @property
    def repartition(self):
        repartition = dict()
        for ingredient in self.ingredients:
            repartition[ingredient.name] = ingredient.quantity
        return repartition

    def __repr__(self):
        return f"{self.repartition}=>{self.score}"


def partial_generator(repartition, ingredients, current_ingredient,
                      pending_quantity):
    ingredient = ingredients[current_ingredient]
    if not current_ingredient:
        repartition[ingredient.name] = pending_quantity
        yield repartition
    else:
        current_ingredient -= 1
        for quantity in range(pending_quantity + 1):
            new_pending_quantity = pending_quantity - quantity
            repartition[ingredient.name] = quantity
            yield from partial_generator(repartition, ingredients,
                                         current_ingredient,
                                         new_pending_quantity)


def repartition_generator(ingredients):
    repartition = dict()
    nb_ingredients = len(ingredients)
    yield from partial_generator(repartition, ingredients, nb_ingredients - 1,
                                 100)


def main():
    ingredients = [
        Ingredient('Butterscoth', -1, -2, 6, 3, 8),
        Ingredient('Cinnamon', 2, 3, -2, -1, 3),
    ]
    recipe = Recipe(ingredients, {'Butterscoth': 44, 'Cinnamon': 56})
    print(recipe)
    with open('input.txt') as fd:
        ingredients = fd.readlines()
    ingredients = (ing.rstrip('\n') for ing in ingredients)
    ingredients = (INGREDIENT_SPECS_RE.match(ing) for ing in ingredients)
    ingredients = filter(lambda x: x, ingredients)
    ingredients = map(lambda ing: ing.group, ingredients)
    ingredients = map(Ingredient.from_spec_re, ingredients)
    ingredients = list(ingredients)
    recipes = (Recipe(ingredients, repartition)
               for repartition in repartition_generator(ingredients))
    print(max(recipes, key=lambda recipe: recipe.score))
    recipes = (Recipe(ingredients, repartition)
               for repartition in repartition_generator(ingredients))
    recipes = (recipe for recipe in recipes if recipe.calories == 500)
    print(max(recipes, key=lambda recipe: recipe.score))


if __name__ == '__main__':
    main()
