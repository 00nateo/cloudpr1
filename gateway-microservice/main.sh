#!/bin/bash
# Start Go server
echo "Starting Go server..."
cd /app
./chatbotservice &
GO_PID=$!

# wait for processes
wait $GO_PID


