#! /bin/bash

nameCours="Programming in bash shell"
scheduleCours="weekend"

if [ $# -ge 1 ]; then
    nameCours=$1
    if [ $# -eq 2 ]; then
        scheduleCours=$2
    fi
fi

echo -e "The name of course is \e[32m $nameCours \e[0m in the schedule \e[32m $scheduleCours \e[0m"
