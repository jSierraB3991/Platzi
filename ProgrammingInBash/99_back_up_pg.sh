#! /bin/bash

option=''
backup_name=''
password=""

echo "Program Utilities for PostrgeSql"
read -n1 -p "Choose Option: " option
echo ""
read -p "Whats name for backuo?: " backup_name
echo "Option $option, BackUpName $backup_name"

read -s -p "Key: " password
echo ""
echo "key: " $password

