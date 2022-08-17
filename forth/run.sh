#!/bin/bash

FILE=${1:-hello.fs}

gforth $FILE -e bye
