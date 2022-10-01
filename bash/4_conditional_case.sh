#! /bin/bash

option=''

read -n1 -p  "Selecction a operation, of a to z: " option
echo ""
case $option in
    "a") echo "Installing PostgresSql";;
    "b") echo "Deleting files";;
    [c-h]) echo "Operation not implement";;
    "i") echo "Reconfiguaration for PostgreSql";;
    [j-z]) echo "Opertation not implemtn";;
    *) echo "Incorrected Option"
esac

