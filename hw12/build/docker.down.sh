#!/bin/bash
docker-compose  --file build/docker/docker-compose.yml down
docker image rm hw12_calendar:latest --force