#!/usr/bin/env python3

import random

MAXGEN = 10000
NONWORD = '\n'

w1 = w2 = NONWORD

with open('Genesis.txt', 'r') as fh:
    data = fh.read()

lines = data.split("\n")

statetab = dict()

for line in lines:
    words = line.split(" ")
    for word in words:
        if not w1 in statetab:
            statetab[w1] = dict()
        if not w2 in statetab[w1]:
            statetab[w1][w2] = list()
        statetab[w1][w2].append(word)
        w1 = w2
        w2 = word

statetab[w1][w2].append(NONWORD)

w1 = w2 = NONWORD
for i in range(0, MAXGEN):
    suf = statetab[w1][w2]
    r = random.randint(0, len(suf)-1)
    t = suf[r]
    if (t == NONWORD):
        exit()
    print(t)
    w1 = w2
    w2 = t
