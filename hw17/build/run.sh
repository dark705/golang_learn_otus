#!/bin/bash
docker-compose -f ./build/docker/docker-compose-environment.yml up -d
sleep 10
docker-compose -f ./build/docker/docker-compose.yml up