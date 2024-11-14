#!/bin/bash

# Invoke the Lambda function with a test event
output=$(echo '{"name": "HealthCheck"}' | ./bootstrap)

# Check if the output contains the expected response
if echo "$output" | grep -q "Hello HealthCheck!"; then
    exit 0
else
    exit 1
fi