#!/bin/bash
docker build --tag hw14_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw14_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./
docker-compose --file build/docker/docker-compose.yml up --detach