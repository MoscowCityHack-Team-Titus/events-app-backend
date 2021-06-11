FROM golang:latest

WORKDIR /events-server

ENV PROJ_ROOT=/events-server

COPY . .

RUN apt-get update && \
    make build-app