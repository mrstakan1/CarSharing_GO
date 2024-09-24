#!/usr/bin/env bash
# scripts/wait-for-it.sh

set -e

TIMEOUT=15
DELAY=1

HOST=$(echo $1 | cut -d':' -f1)
PORT=$(echo $1 | cut -d':' -f2)

shift

COMMAND="$@"

for ((i=0;i<$TIMEOUT;i++)); do
    if nc -z $HOST $PORT; then
        echo "Service is up - executing command"
        exec $COMMAND
    fi
    echo "Waiting for $HOST:$PORT..."
    sleep $DELAY
done

echo "Service $HOST:$PORT not available after $TIMEOUT seconds"
exit 1
