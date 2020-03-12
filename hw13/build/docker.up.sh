#!/bin/bash
docker build --tag hw13_calendar:latest -f ./build/docker/Dockerfile ./
docker-compose --file build/docker/docker-compose.yml up --detach