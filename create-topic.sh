#!/bin/sh

echo "Waiting for Kafka..."
sleep 5

echo "Checking if the topic 'notifications' exists..."
if kafka-topics --list --bootstrap-server kafka:9092 | grep -q "^notifications$"; then
  echo "The topic 'notifications' already exists. No need to create it."
else
  echo "Creating the topic 'notifications'..."
  kafka-topics --create --topic notifications --bootstrap-server kafka:9092 --partitions 3 --replication-factor 1
fi

exit 0
