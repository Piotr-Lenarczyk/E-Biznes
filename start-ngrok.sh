#!/bin/bash

# Start Play Framework application in the background
sbt run &

# Wait for the server to start
echo "Waiting for Play server to be ready..."
until [[ "$(curl -s -o /dev/null -w '%{http_code}' http://localhost:9000/)" == "200" ]]; do
  sleep 5
done
echo "Play server is up and running!"

echo "Play app is running, starting ngrok..."

# Authenticate ngrok using the token from environment variables
ngrok authtoken "$NGROK_AUTH_TOKEN"

# Kill any existing ngrok sessions (if running)
pkill -f ngrok || true

# Start ngrok and expose the Play app on port 9000
ngrok http 9000 > /dev/null &  # Run ngrok in the background
sleep 20  # Give it some time to establish the tunnel

# Get and print the public URL
NGROK_URL=$(curl -s http://127.0.0.1:4040/api/tunnels | jq -r '.tunnels[0].public_url')

if [ -n "$NGROK_URL" ]; then
    echo "Ngrok tunnel is live! Public URL: $NGROK_URL"
else
    echo "Failed to get ngrok URL."
    exit 1
fi

# Keep the container running
wait
