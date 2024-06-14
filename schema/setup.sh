#!/bin/sh
# update the config properly before run

dbname="txbank"
port=${PGPORT:-5656}
user="txbank"
password="txbank"
host=${PGHOST:-localhost}

PGPASSWORD=$password psql -U $user -d "postgres" -h $host -p $port -c "CREATE DATABASE $dbname" 2> /dev/null
for filename in schema/*.sql; do
    PGPASSWORD=$password psql -h $host -p $port -d $dbname -U $user -f "$filename"
done
