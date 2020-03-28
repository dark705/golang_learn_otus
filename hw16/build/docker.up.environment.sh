#!/bin/bash
docker build --tag hw16_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw16_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./
docker-compose --file build/docker/docker-compose-environment.yml up --detach