#!/bin/bash
docker image rm hw16_postgres:latest --force
docker image rm hw16_rabbitmq:latest --force

docker image rm hw16_calendar_scheduler:latest --force
docker image rm hw16_calendar_sender:latest --force
docker image rm hw16_calendar_api:latest --force
docker image rm hw16_calendar_godog:latest --force