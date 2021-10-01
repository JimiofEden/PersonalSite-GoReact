#!/bin/bash

file="/docker-entrypoint-initdb.d/DatabaseDump.pgdata"
dbname=202109PersonalSite

echo "Restoring DB using $file"
pg_restore -U postgres --dbname=$dbname --verbose --single-transaction < "$file" || exit 1