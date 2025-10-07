#!/bin/bash

find . -type f -name '*.sh' -printf '%f\n' |      # find files ending with .sh, print just filename
sed 's/\.sh$//' |                                 # remove the .sh extension
sort -r                                          # sort in descending order
