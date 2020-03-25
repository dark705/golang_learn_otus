#!/bin/bash
docker build --tag hw15_postgres:latest -f ./build/docker/postgres/Dockerfile ./
docker build --tag hw15_rabbitmq:latest -f ./build/docker/rabbitmq/Dockerfile ./

docker build --tag hw15_calendar_scheduler:latest -f ./build/docker/calendar/scheduler/Dockerfile ./
docker build --tag hw15_calendar_sender:latest -f ./build/docker/calendar/sender/Dockerfile ./
docker build --tag hw15_calendar_api:latest -f ./build/docker/calendar/api/Dockerfile ./
