#!/bin/bash
docker build --tag hw15_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw15_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./
docker-compose --file build/docker/docker-compose.yml up --detach