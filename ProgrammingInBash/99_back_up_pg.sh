#! /bin/bash

option=''

function install_postgre() {
    echo -e "\nInstalling postgres....."
}

function unistall_postgre() {
    echo -e "\nUninstalling postgres...."
}

function do_backup() {
    echo -e "\nDoing Backup..."
}
function show_backup() {
    echo -e "\nShowing Backup..."
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

        3) do_backup;;

        4) show_backup;;

        5)  echo -e "\nExit Program"
            exit 0;;
    esac
    sleep 3
done    

