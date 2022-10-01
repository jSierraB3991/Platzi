#! /bin/bash

URL="http://localhost:5000/schedule/dynamic"
METHOD="-X POST"
HEADER="-H \"content-type: application/json\""
BODY=""

if [ "$HEADER" != "" ]; then
    continue="s"
    BODY="-d \"{"
    while [ "$continue" == "s" ]; do
        read -p "What is a key for body: " key
        if [ "$key" != "" ]; then
            read -p "What is a value for key $key: " value
            if [ "$value" != "" ]; then
                BODY=$(echo $BODY "\\\"$key\\\": \\\"$value\\\"" )
            else
                echo "invalid value for $key, no save"
            fi
        else
            echo "invalid key, no save"
        fi

        read -p "Continue add body in request: S/n " continue
        if [ "$continue" == "" ] || [ "$continue" == "S" ]; then
            continue="s"
            BODY="$BODY, "
        fi
    done
    BODY=$(echo $BODY "}\"")
fi
REQUEST=$(echo "curl $METHOD $HEADER $BODY $URL")
echo $REQUEST | xclip -sel clip
echo "$REQUEST al cortapapeles"

