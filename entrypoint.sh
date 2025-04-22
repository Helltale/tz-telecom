#!/usr/bin/env bash
set -e

echo "waiting for PostgreSQL to be available..."
./wait-for-it.sh tz-postgres 5432 -- ./tz-telecom migrate

echo "complete, starting application..."
exec ./tz-telecom serve
