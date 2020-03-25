#!/bin/bash
docker-compose  --file build/docker/docker-compose.yml down
docker image rm hw15_postgres:latest --force
docker image rm hw15_rabbitmq:latest --force