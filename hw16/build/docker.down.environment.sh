#!/bin/bash
docker-compose  --file build/docker/docker-compose-environment.yml down
docker image rm hw16_postgres:latest --force
docker image rm hw16_rabbitmq:latest --force