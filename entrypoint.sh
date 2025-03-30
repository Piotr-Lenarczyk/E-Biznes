#!/bin/bash

# Start Play Framework application in the background
sbt run &

# Wait for the server to start
echo "Waiting for Play server to be ready..."
until [[ "$(curl -s -o /dev/null -w '%{http_code}' http://localhost:9000/)" == "200" ]]; do
  sleep 5
done
echo "Play server is up and running!"

# Run API tests
./test_api.sh

# Stop the Play application after tests
echo "Stopping Play server..."
pkill -f sbt
