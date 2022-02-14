#! /bin/bash

option=''

while :
do
    clear
    echo "_________________________________________"
    echo "PGUTIL - Program Utility for Postgre Sql "
    echo "_________________________________________"
    echo "                  MAIN MENU              "
    echo "_________________________________________"
    echo "1. Install Postgres"
    echo "2. UnInstall Postgres"
    echo "3. Do Backup"
    echo "4. Show Backup"
    echo "5. exit"

    read -n1 -p "Choose the option [1-5]: " opcion

    case $opcion in
        1)
            echo -e "\nInstalling postgres....."
            sleep 3;;
        2) 
            echo -e "\nUninstalling postgres...."
            sleep 3;;
        3) 
            echo -e "\nDoing Backup..."
            sleep 3;;
        4) 
            echo -e "\nShowing Backup..."
            sleep 3;;
        5)  
            echo -e "\nExit Program"
            exit 0;;
    esac
done    

