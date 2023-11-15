#!/bin/sh

if [[ $# -ne 2 ]]; then
    echo "Useage: create-user.sh \"username\" \"password\""
    exit
fi

echo "$2" | sha256sum | sed 's/...$//' > users/$1.txt
mkdir storage/$1
