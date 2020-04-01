#!/bin/bash
docker-compose  --file build/docker/docker-compose-tests.yml down

#docker image rm hw16_postgres:latest --force
#docker image rm hw16_rabbitmq:latest --force

#docker image rm hw16_calendar_scheduler:latest -f --force
#docker image rm hw16_calendar_sender:latest -f --force
#docker image rm hw16_calendar_api:latest -f --force
docker image rm hw16_calendar_godog:latest -f --force

