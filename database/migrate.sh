#!/bin/sh

migrate -path /database/migrations -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_ADDR)/$MYSQL_DATABASE" --verbose up