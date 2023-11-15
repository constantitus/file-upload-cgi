#!/bin/sh

if [[ $# -ne 2 ]]; then
    echo "Useage: create-user.sh \"username\" \"password\""
    exit
fi

echo "$2" | sha256sum | sed 's/...$//' > $1.txt
