#!/bin/bash
docker-compose  --file build/docker/docker-compose.yml down
docker image rm hw14_postgres:latest --force
docker image rm hw14_rabbitmq:latest --force