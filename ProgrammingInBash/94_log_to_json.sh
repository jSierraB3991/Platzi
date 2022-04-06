#! /bin/bash

set -e

if [ ! $# -eq 1 ]; then
    echo -e "\n \e[31m I NEED for parameter the file \e[0m"
    exit 1
fi

file=$1
if [ ! -f $file ]; then
    echo -e "\n \e[31m The file $1, not exists \e[0m"
    exit 1
fi

STRING_EXIT="{"
file2=$(echo "/home/juan-sierra/file2.txt")
if [ ! -f $file2 ]; then
    touch $file2
fi

/usr/bin/cat $file | sed 's/\,/\n/g' | sed 's/=/:/g' | sed 's/(/\n {/g' | sed 's/)/\n }/g' >  $file2

while read -r line;
do
    isObject=$(echo $line | awk 'BEGIN{FS=":"} {print $2}')
    isObject2=$(echo $line | awk 'BEGIN{FS=":"} {print $1}')

    if [ "$(echo $isObject | grep "^\[" )" != "" ]; then
        objectname=$(echo $line | awk 'BEGIN{FS=":"} {print $1}')
        echo "\"$objectname\" : [{"
    elif [ "$(echo $isObject2 | grep "^{" )" != "" ];then
        objectname=$(echo $line | awk 'BEGIN{FS=":"} {print $1}' | sed 's/{/ /g')
        echo "\"$objectname\" : {"
    elif [ "$line" == "}" ]; then
        echo "$line, "
    else
        attributte=$(echo $line | sed 's/:/\": \"/g')
        echo "\"$attributte\","
    fi
done < $file2

echo -e "\e[32m Parsed Sucessfull \e[0m"
echo "$STRING_EXIT"
echo "}"
