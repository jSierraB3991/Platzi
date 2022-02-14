#! /bin/bash

arreglo_numbers=(1 2 3 4 5 6 7 8 9 0)
arreglo_strings=(Ana, Pedro, Marco, Antonio)
arreglo_ranges=({A..Z} {0..10})

echo -e "\e[32mArrays of number: \e[31m${arreglo_numbers[*]}"
echo -e "\e[32mArrays of strings: \e[31m${arreglo_strings[*]}"
echo -e "\e[32mArrays of ranges: \e[31m${arreglo_ranges[*]}"

echo -e "\e[32mLength of array number \e[31m${#arreglo_numbers[*]}"
echo -e "\e[32mLength of array strings \e[31m${#arreglo_strings[*]}"
echo -e "\e[32mLength of array ranges \e[31m${#arreglo_ranges[*]}"

echo -e "\e[32mPosition index 3 for array number \e[31m${arreglo_numbers[3]}"
