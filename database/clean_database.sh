#!/bin/sh


if [ "$ENV" = "dev" ] || [ "$ENV" = "test" ]; then
  migrate -path /database/migrations -database "mysql://$MYSQL_USER:$MYSQL_PASSWORD@tcp($MYSQL_ADDR)/$MYSQL_DATABASE" drop -f
fi