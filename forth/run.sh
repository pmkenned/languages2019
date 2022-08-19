#!/bin/bash

FILE=${1:-fib.fs}

gforth $FILE -e bye
