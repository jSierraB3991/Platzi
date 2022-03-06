#! /bin/bash

option=''

install_on_debian_based() {
    sudo apt update
    if [ "$(hostnamectl | grep Operating | awk '{print $3}')" == "Parrot" ]; then
	sudo parrot-upgrade
    else
	sudo apt upgrade -y
    fi

    sudo apt install postgresql postgresql-contrib -y
}

install_on_arch_based() {
    sudo pacman -Syu --noconfirm
    sudo pacman -S postgresql --noconfirm
    sudo -i -u postgres initdb -D '/var/lib/postgres/data'
}

function enable_postgre_service_systemd() {
    sudo systemctl enable postgresql.service
}

function enable_postgre_service_open_rc() {
    sudo rc-update add postgresql default
}

function decide_init_and_run_function() {
    function_systemd=$1
    function_open_rc=$2

    init=$(cat /proc/1/comm)
    if [ "$init" == "systemd" ]; then
        $function_systemd
    elif [ "$init" == "openrc" ] || [ "$init" == "openrc-init" ]
        $function_open_rc
    fi
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
            decide_init_and_run_function run_postgre_service_systemd run_postgre_service_open_rc
        fi
        
    fi
    echo -e "\n"
    read -n1 -p "press any key to continue"
}

unistall_in_debian_based() {
    sudo apt-get -y --purge remove postgresql\*
}

unistall_in_arch_based() {
    sudo pacman -Rsn postgresql --noconfirm
}

function stop_and_disable_postgre_service_systemd() {
    sudo systemctl stop postgresql.service
    sudo systemctl disable postgresql.service
}

function stp_and_disable_postgre_service_open_rc() {
    sudo rc-service postgresql stop
    sudo rc-update del postgresql
}

function unistall_postgre() {
    echo -e "\nVerifing Postgre"
    if [ "$(which psql)" == "" ]; then
        echo "Postgres not is install"
    else
	echo -e "\nUninstalling postgres...."
        decide_init_and_run_function stop_and_disable_postgre_service_systemd stp_and_disable_postgre_service_open_rc

        if [ "$(which apt)" != "" ]; then
            unistall_in_debian_based
        elif [ "$(which pacman)" != "" ]; then
            unistall_in_arch_based
        fi

        sudo rm -r /var/lib/postgres
        sudo userdel -r postgres
        sudo groupdel postgresql
    fi
    echo -e "\n"
    read -n1 -p "press any key to continue"
}

function verify_status_postgresql_with_system() {
    init=$(cat /proc/1/comm)
    if [ "$init" == "systemd" ]; then
        echo "systemctl status postgresql.service"
    elif [ "$init" == "openrc" ] || [ "$init" == "openrc-init" ]
        echo "rc-service postgresql status"
    fi
}

function verify_service() {

    if [ "$(which psql)" != "" ]; then
        state_service=$($(verify_status_postgresql_with_system) | grep Active | awk '{print $2}')

	if [ "$state_service" == "inactive" ]; then
	    echo "You have active postgresql service"
	    return 0
	fi
	return 1
    else
	echo "You don't have installing postgresql"
	return 0
    fi
    return 0
}

function do_backup() {
    verify_service
    if [ "$?" == "1" ]; then

        read -p "directory of backup: " back_up_directory
	if [ "$(which psql)" != "" ]; then
	    back_up_directory=$1

            cd /
	    sudo -u postgres psql -c "\l"

	    read -p "Choose database to backup: " db_backup
        
            if [ ! -d $back_up_directory ]; then
	        mkdir -m 755 $back_up_directory
	    fi
        
            if [ -d $back_up_directory ]; then
	        echo "Doing Backup... $back_up_directory"
	        name_backup="db_backup_$(date +%s).bak"
            
		sudo -u postgres pg_dump -Fc $db_backup > "$back_up_directory/$name_backup"
                echo "permiss of directory"
	        sudo chmod 755 $back_up_directory
	        echo "Backup Ok in $back_up_directory/$name_backup"

		cd -
            else
	        echo -e "The directory $back_up_directory not exists"
	    fi

        else
	    echo -e "Postgresql is not install"
        fi
    fi
    read -n1 -p "press any key to continue"
    echo -e "\n"
}
function show_backup() {
    verify_service
    if [ "$?" == "1"  ]; then
	read -p "directory of backup: " back_up_directory

	if [ -d $back_up_directory ]; then

            echo "backup list"
	    ls $back_up_directory | grep "**.bak"
	    read -p "Choose backup to backuping: " backuping

            if [ -f $back_up_directory/$backuping ]; then
	        cd / 

                sudo -u postgres psql -c "\l"
	        read -p "name of data base destiny: " db_destiny

	        db_name=$(sudo -u postgres psql -lqt | grep "$db_destiny" | awk '{print $1}' )
		if [ "$db_name" != "$db_destiny" ]; then
		    sudo -u postgres psql -c "CREATE DATABASE $db_destiny"
                fi

                echo "backuping in the data base $db_destiny"
	        sudo -u postgres pg_restore -Fc -d $db_destiny "$back_up_directory/$backuping"
	        echo "Lis of database names"
		sudo -u postgres psql -c "\l"
            
                cd -
	    else
	       echo "The backup $backuping not exists"
            fi
	else
	    echo "The directory $back_up_directory not exists"
        fi
    fi
    read -n1 -p "press any key to continue"
    echo -e "\n"
}

function start_service_systemd() {
    sudo systemctl start postgresql.service
}

function start_service_openrc() {
    sudo rc-service start postgresql
}

function start_service() {

    if [ "$(which psql)" != "" ]; then
        state_service=$($(verify_status_postgresql_with_system) | grep Active | awk '{print $2}')

	if [ "$state_service" == "inactive" ]; then
	    echo "Starting service of postgresql"
            decide_init_and_run_function start_service_systemd start_service_openrc
	else
	    echo "The Service postgresql in active"
        fi
    else 
	echo "You not installing postgresql"
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
    echo "5. Start service of postgre"
    echo "6. exit"

    read -n1 -p "Choose the option [1-6]: " opcion

    case $opcion in
        1) install_postgre;;

        2) unistall_postgre;;

        3) echo -e "\n"
	    do_backup;;

        4) echo -e "\n"
	    show_backup;;

        5) echo -e "\n"
	    start_service;;

        6)  echo -e "\nExit Program"
            exit 0;;
    esac
done    

