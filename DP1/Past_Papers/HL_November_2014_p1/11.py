#!/usr/bin/env python3
from lib.functions import createMagicSquare, pprint

s = [
        [30, 39, 48, 1, 10, 19, 28],
        [38, 47, 7, 9, 18, 27, 29],
        [46, 6, 8, 17, 26, 35, 37],
        [5, 14, 16, 25, 34, 36, 45],
        [13, 15, 24, 33, 42, 44,4],
        [21, 23, 32, 41, 43, 3, 12],
        [22, 31, 40, 49, 2, 11, 20],
]

n = 5

pprint(createMagicSquare(n))