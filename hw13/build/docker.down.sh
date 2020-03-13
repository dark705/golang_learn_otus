#!/bin/bash
docker-compose  --file build/docker/docker-compose.yml down
docker image rm hw13_calendar:latest --force