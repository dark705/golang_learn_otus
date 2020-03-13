#!/bin/bash
docker build --tag hw13_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw13_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./
docker-compose --file build/docker/docker-compose.yml up --detach