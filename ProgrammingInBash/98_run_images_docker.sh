#! /bin/bash

folder=$(echo "${PWD}")

function what_container() {
    if [ "$(which docker)" != "" ]; then
        echo "docker"
    elif [ "$(which podman)" != "" ]; then
        echo "podman"
    fi
}

function verify_run_container(){
    container_name=$1
    container_provider=$(what_container)
    echo -e "\e[32mVerifing if container $container_name is running\e[0m"
    container_run=$(sudo $container_provider ps --format "{{.Names}}" | grep $container_name)

    if [ "$container_run" != "" ]; then
        return 1
    fi
    return 0
}

function remove_stop_container(){
    container_name=$1
    container_provider=$(what_container)
    container_run=$(sudo $container_provider ps -a --format "{{.Names}}" | grep $container_name)

    if [ "$container_run" != "" ]; then
        id_container=$( \
            sudo $container_provider ps --format "{{.ID}}\t.{{.Names}}" \
                | grep $container | awk '{print $1}' )
        echo -e "[31mKilling container $container with id $id_container"
        sudo $container_provider rm $id_container
    fi
}

function verify_container() {
    container=$1
    function=$2

    verify_run_container $container
    if [ "$?" == "1" ]; then
        read -n1 -p "Already container $container is run, do you like kill?: Y/n " kill_container
        echo ""

        if [ "$kill_container" == "y" ] || [ "$kill_container" == "Y" ] || 
            [ "$kill_container" == "" ]; then

            echo -e "\e[32mKilling container $container\e[0m"
            container_provider=$(what_container)
            id_container="$( \
                sudo $container_provider ps --format "{{.ID}}\t.{{.Names}}" \
                    | grep $container | awk '{print $1}' )"

            sudo $container_provider stop $id_container
            
            already_exists="$id_container"
            while [ "$already_exists" == "$id_container" ]; do
               already_exists="$( \
                sudo $container_provider ps -a --format "{{.ID}}\t.{{.Names}}" \
                    | grep $container | awk '{print $1}')"
            done
            $function
        else
            echo -e "\e[31mCanceling run of container $container\e[0m"
        fi
    else
        remove_stop_container $container
        $function
    fi
}

function what_image() {
    container_image=$1
    if [ "$(which docker)" != "" ]; then
        echo "$container_image"
    elif [ "$(which podman)" != "" ]; then
        echo "docker.io/library/$container_image"
    fi
}

function run_mongo_inscription() {
    echo -e "\e[32mRUN CONTAINER mongo-inscription\e[0m"
    sudo $(what_container) run --rm -d -p 27017:27017 --name mongo-inscription $(what_image mongo):5.0.3-focal
}

function run-postgre-database() {
    container_provider=$(what_container)
    port=5432
    name=postgre_zabud
    array_db_names=(zabud_inscription zabud_core zabud_notification zabud_planning)
    tiemp_of_sleep=4

    volumes=" -v $ZABUD_HOME/data/$name:/var/lib/postgresql/data"
    enviorment=" -e POSTGRES_USER=postgres -e POSTGRES_DB=postgres -e POSTGRES_PASSWORD=root"
    configurations=" run --rm -d -p $port:5432 --name $name $volumes $enviorment"

    echo -e "\e[32mRUN CONTAINER $name\e[0m"
    sudo $container_provider $configurations postgres:12.9-alpine

    echo "Sleeping $tiemp_of_sleep seconds"
    sleep $tiemp_of_sleep
    echo "Wake up"

    for db_name in ${array_db_names[*]}; do
        db_exists=$(sudo $container_provider exec $name psql -U postgres -lqt | grep $db_name | awk '{print $1}')
        if [ "$db_exists" == "" ]; then
            echo -e "\e[32mCreating Datbase $db_name\e[0m"
            sudo $container_provider exec $name psql -U postgres -c "CREATE DATABASE $db_name"
        else
            echo "The database $db_name already exists"
        fi
    done
}

function pg_docker_dbs() {
    verify_container postgre_zabud run-postgre-database 
}

function queue_activemq() {
    echo -e "\e[32mRUN CONTAINER activemq\e[0m"
    sudo $(what_container) run --rm --name activemq -d -p 8161:8161 -p 61616:61616 rmohr/activemq:5.14.0-alpine
}

function zookeeper_kafka() {

    container_provider=$(what_container)
    script_container="sudo $container_provider ps --format {{.ID}}\t{{.Names}}"
    id_container_of_zookeper=$($script_container | grep zookeeper | awk '{print $1}' )
    id_container_of_kafka=$($script_container | grep kafka | awk '{print $1}' )

    echo -e "\e[32mVerifing container zookeeper\e[0m"
    if [ "$id_container_of_zookeper" == "" ]; then
    echo -e "\e[32mRUN CONTAINER Zookeeper\e[0m"
        sudo $container_provider run --rm --name zookeeper -d -p 2181:2181 wurstmeister/zookeeper
    else
        echo "The container of Zookeeper zookeeper is already exists"
    fi

    echo -e "\e[32mVerifing container Kafka\e[0m"
    if [ "$id_container_of_kafka" == "" ]; then
        ip_private=$(ifconfig wlan0 | grep inet | head -1 | awk '{print $2}')
        enviorment=" -e KAFKA_ADVERTISED_HOST_NAME=localhost"
        enviorment="$enviorment -e KAFKA_ZOOKEEPER_CONNECT=$ip_private:2181"
        configurations="--rm --name kafka -d -p 9092:9092 $enviorment"
        echo -e "\e[32mRUN CONTAINER Kafka\e[0m"
        sudo $container_provider run $configurations wurstmeister/kafka
    else
        echo "The container of Kafka kafka is already exists"
    fi
}

function zabud_discovery() {
    cd $ZABUD_HOME/zabud-discovery-ms
    if [ ! -f $ZABUD_HOME/zabud-discovery-ms/Dockerfile ]; then
        cp $DOT_FILES/Docker/spring-Dockerfile ./Dockerfile
    fi
    container_provider=$(what_container)

    application="zabud-discovery"
    version="1.0"
    echo -e "\e[32mSearch Image $application:$version \e[0m"
    dockerimage=$(sudo $container_provider images --format "{{.Repository}}" $application:$version)
    if [ "$dockerimage" == "" ]; then
        read -p "you image $application:$version not exists, Dou you like build image?: " response

        if [ $response == "y" ] || 
           [ $response == "s" ] || 
           [ $response == "Y" ] || 
           [ $response == "S" ]; then
            cd $ZABUD_HOME/zabud-discovery-ms
            echo -e "\e[32mGenerate Image zabud-discovery\e[0m"
            sudo $container_provider build -t zabud-discovery:1.0 .
        fi
    fi
    dockerimage=$(sudo $container_provider images --format "{{.Repository}}" $application:$version)
    if [ "$dockerimage" != "" ]; then
        echo -e "\e[32mRUN CONTAINER $application\e[0m"
        sudo $container_provider run --rm -d -p 8761:8761 --name $application $application:$version
    fi
    cd $folder
}

function run_help() {
    echo -e "\nrun_zabud_images.sh [-r [OPTIONS]]" \
        "\noptions | containers configurate:" \
        "\n\tmongo_inscription" \
        "\n\tpostgre_zabud" \
        "\n\tqueue_activemq" \
        "\n\tzookeeper_kafka | Zookeeper and Kafka" \
        "\n\tzabud_discovery"
}

function error_to_help() {
    echo -e "\n \e[31mERROR $1\e[0m"
    run_help
}

if [ -z $ZABUD_HOME ]; then
    echo "I Need enviroment varible ZABUD_HOME"
elif [ -z $DOT_FILES ]; then
    echo "I Need enviroment varible DOT_FILES"
else
    if [ $# -eq 0 ]; then
        verify_container mongo-inscription run_mongo_inscription
        pg_docker_dbs
        verify_container activemq queue_activemq
        zookeeper_kafka
        verify_container zabud-discovery zabud_discovery

    elif [ $# -eq 1 ]; then
        if [ "$1" == "-h" ] || [ "$1" == "--help" ]; then
            run_help
        else
            error_to_help "The paramater $1 not exists"
        fi
    elif [ $# -eq 2 ] && [ "$1" == "-r" ]; then
        case $2 in
            "mongo_inscription") verify_container mongo-inscription run_mongo_inscription;;
            "postgre_zabud") pg_docker_dbs;;
            "queue_activemq") verify_container activemq queue_activemq;;
            "zookeeper_kafka") zookeeper_kafka;;
            "zabud_discovery") verify_container zabud-discovery zabud_discovery;;
            *)  error_to_help "The container $2 not configurate";;
        esac

    else
        error_to_help "Acction Failed $1 $2"
    fi
fi
