#! /bin/bash

arreglo_numbers=(1 2 3 4 5 6 7 8 9 0)

echo "Array of numbers"
for num in ${arreglo_numbers[*]}; do
    echo "number:$num"
done

echo -e "\nValues Strings"
for name in "Marco" "Juan" "Pedro" "Daniela"; do
    echo "Name: $name"
done

echo -e "\nFiles"
for file in *; do
    echo "File: $file"
done

echo -e "\nCommand"
for file in $(cat file.txt); do
    echo "File: $file"
done

echo -e "\nTraditional (C)"
for ((i=1; i<10; i++)); do
    echo "number $i"
done
