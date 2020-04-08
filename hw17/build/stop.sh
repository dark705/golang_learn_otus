#!/bin/bash
docker-compose -f ./build/docker/docker-compose.yml down
docker-compose -f ./build/docker/docker-compose-environment.yml down