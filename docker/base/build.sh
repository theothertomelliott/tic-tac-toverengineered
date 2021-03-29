#!/bin/bash

cd "$(dirname "$0")"
docker build -t "docker.io/tictactoverengineered/base:latest" .