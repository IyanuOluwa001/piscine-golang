#!/bin/bash

ls -1 | sort -V | xargs ls -l | sed -n 'p;n'
