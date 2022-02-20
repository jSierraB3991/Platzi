#! /bin/bash

option=''

install_on_debian_based() {
    sudo apt update
    sudo apt upgrade -y
    sudo apt install postgresql postgresql-contrib
}

install_on_arch_based() {
    sudo pacman -Syu --noconfirm
    sudo pacman -S postgresql --noconfirm
    sudo -i -u postgres initdb -D '/var/lib/postgres/data'
}

function install_postgre() {

    echo -e "\nVerifing Postgre"
    if [ "$(which psql)" != "" ]; then
        echo "Postgres is install"
    else
        read -s -p "password for postgre " postgre_password
        
        echo -e "\nUpdate System and Installing postgres....."
        if [ "$(whoami)" != "root" ]; then
            echo "Need auth with super user for installing postgresql"
        fi

        if [ "$(which apt)" != "" ]; then
            install_on_debian_based
        elif [ "$(which pacman)" != "" ]; then
            install_on_arch_based
        fi

        sudo postgres psql -c "ALTER USER postgres WITH PASSWORD '{$postgre_password}'"
        
        read -n1 -p "¿You want enable postgre service [Y/n] ?" enable_postgre
        if [ "$enable_postgre" == "Y" ] || [ "$enable_postgre" == "y" ] || [ "$enable_postgre" == "" ]; then
            sudo systemctl enable postgresql.service
        fi
        
    fi
    echo -e "\n"
    read -n1 -p "press any key to continue"
}

unistall_in_debian_based() {
    sudo apt-get -y --purge remove postgresql\*
}

unistall_in_arch_based() {
    sudo pacman -R postgresql
    sudo pacman -R postgresql-libs
}

function unistall_postgre() {
    echo -e "\nUninstalling postgres...."
    sudo systemctl stop postgresql.service
    
    if [ "$(which apt)" != "" ]; then
        unistall_in_debian_based
    elif [ "$(which pacman)" != "" ]; then
        install_on_arch_based
    fi

    sudo rm -r /var/lib/postgres
    sudo userdel -r postgres
    sudo groupdel postgresql
    echo -e "\n"
    read -n1 -p "press any key to continue"
}

function do_backup() {
    if [ "$(which psql)" != "" ]; then
        back_up_directory=$1

        cd /
        sudo -u postgres psql -c "\l"
        cd -

        read -p "Choose database to backup: " db_backup
        
        if [ ! -d $back_up_directory ]; then
            mkdir $back_up_directory
        fi
        
        if [ -d $back_up_directory ]; then
            echo "Doing Backup... $back_up_directory"
            name_backup="db_backup_$(date +%s).bak"
            cd /
            sudo -u postgres pg_dump -Fc $db_backup > "$back_up_directory/$name_backup"
            cd -
            
            echo "permiss of directory"
            sudo chmod 755 $back_up_directory
            echo "Backup Ok in $back_up_directory/$name_backup"
        else
            echo -e "The directory $back_up_directory not exists"
        fi

    else
        echo -e "Postgresql is not install"
    fi

    read -n1 -p "press any key to continue"
    echo -e "\n"
}
function show_backup() {
    back_up_directory=$1

    if [ -d $back_up_directory ]; then

        echo "backup list"
        ls $back_up_directory | grep "**.bak"
        read -p "Choose backup to backuping: " backuping

        if [ -f $back_up_directory/$backuping ]; then
            cd /
            sudo -u postgres psql -c "\l"
            cd -
            read -p "name of data base destiny: " db_destiny

            db_name=$(cd / && sudo -u postgres psql -lqt | grep "$db_destiny" | awk '{print $1}' )
            cd -
            if [ "$db_name" != "$db_destiny" ]; then
                sudo -u postgres psql -c "CREATE DATABASE $db_destiny"
            fi

            echo "backuping in the data base $db_destiny"
            sudo -u postgres pg_restore -Fc -d $db_destiny "$back_up_directory/$backuping"
            echo "Lis of database names"
            sudo -u postgres psql -c "\l"
        # echo "Showing Backup... $back_up_directory"
        else
           echo "The backup $backuping not exists"
        fi
    else
        echo "The directory $back_up_directory not exists"
    fi
    read -n1 -p "press any key to continue"
    echo -e "\n"
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
done    

