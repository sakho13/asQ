# !/bin/bash

# until mysqladmin ping -h db -P 3310 --silent; do
#   echo 'waiting for db to be connectable...'
#   sleep 5
# done

# echo "app is starting...!"
exec go run main.go