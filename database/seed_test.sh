#!/bin/sh

mysql --host="$MYSQL_ADDR" --user="$MYSQL_USER" --password="$MYSQL_PASSWORD" --database="$MYSQL_DATABASE" < /database/seed_test.sql