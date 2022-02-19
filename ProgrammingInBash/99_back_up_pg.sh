#! /bin/bash

option=''

function install_postgre() {
    echo -e "\nInstalling postgres....."
}

function unistall_postgre() {
    echo -e "\nUninstalling postgres...."
}

function do_backup() {
    back_up_directory=$1
    echo "Doing Backup... $back_up_directory"
}
function show_backup() {
    back_up_directory=$1
    echo "Showing Backup... $back_up_directory"
}

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
        1) install_postgre;;

        2) unistall_postgre;;

        3) echo -e "\n"
            read -p "directory of backup: " back_up_directory
            do_backup $back_up_directory;;

        4) echo -e "\n"
            read -p "directory of backup: " back_up_directory
            show_backup $back_up_directory;;

        5)  echo -e "\nExit Program"
            exit 0;;
    esac
    sleep 3
done    

