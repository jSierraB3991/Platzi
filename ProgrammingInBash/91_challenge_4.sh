#! /bin/bash

option=''

while :
do
    clear
    echo "1. Process"
    echo "2  Memory".
    echo "3. Disk Free"
    echo "4. Network Information"
    echo "5. Variable Enviroment".
    echo "6. Program Information"
    echo "7. Backup Information"
    echo "8. exit"
    read -n1 -p "Choose the option [1-8]: " opcion

    case $opcion in
        1)
            echo -e "\n Show Proccess...."
            sleep 3;;
        2) 
            echo -e "\n Show Memory Ram Information...."
            sleep 3;;
        3) 
            echo -e "\n Show Disk Information..."
            sleep 3;;
        4) 
            echo -e "\n Show Network Information..."
            sleep 3;;
        5) 
            echo -e "\n Show Variable Enviroment..."
            sleep 3;;
        6) 
            echo -e "\n Show Progrma Information..."
            sleep 3;;
        7) 
            echo -e "\n Show Information on Backup..."
            sleep 3;;
        8)  
            echo -e "\nExit Program"
            exit 0;;
    esac
done    

