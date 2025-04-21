#!/usr/bin/env bash
set -e

host="$1"
port="$2"
shift 2
cmd="$@"

until nc -z "$host" "$port"; do
  echo "waiting for $host:$port to be available..."
  sleep 1
done

exec $cmd
