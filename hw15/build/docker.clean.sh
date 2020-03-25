#!/bin/bash
docker image rm hw15_postgres:latest --force
docker image rm hw15_rabbitmq:latest --force

docker image rm hw15_calendar_scheduler:latest --force
docker image rm hw15_calendar_sender:latest --force
docker image rm hw15_calendar_api:latest --force