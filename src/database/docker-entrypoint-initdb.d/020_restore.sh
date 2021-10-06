#!/bin/bash

file="/docker-entrypoint-initdb.d/db.dump"
dbname=202109PersonalSite

echo "Restoring DB using $file"
pg_restore -U jimi --dbname=$dbname --verbose --single-transaction < "$file" || exit 1