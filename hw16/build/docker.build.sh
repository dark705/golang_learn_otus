#!/bin/bash
docker build --tag hw16_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw16_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./

docker build --tag hw16_calendar_scheduler:latest -f ./build/docker/calendar/scheduler/Dockerfile ./
docker build --tag hw16_calendar_sender:latest -f ./build/docker/calendar/sender/Dockerfile ./
docker build --tag hw16_calendar_api:latest -f ./build/docker/calendar/api/Dockerfile ./
