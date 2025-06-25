#!/bin/bash

# Health check script for Render backend
# This script pings the health endpoint to keep the service awake

HEALTH_URL="https://dating-backend-wzzl.onrender.com/ping"
LOG_FILE="/tmp/health-check.log"

echo "$(date): Starting health check..." >> $LOG_FILE

# Ping the health endpoint
response=$(curl -s -w "%{http_code}" -o /dev/null $HEALTH_URL)

if [ $response -eq 200 ]; then
    echo "$(date): Health check successful - Status: $response" >> $LOG_FILE
    exit 0
else
    echo "$(date): Health check failed - Status: $response" >> $LOG_FILE
    exit 1
fi 