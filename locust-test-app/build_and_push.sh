#!/usr/bin/env bash

docker build -t locust-test-app:latest .
docker tag locust-test-app:latest yatzek/locust-test-app:latest
docker push yatzek/locust-test-app:latest

