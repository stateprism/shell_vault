#!/usr/bin/env bash

make server

cd server || exit 1

export PRISMA_CA_ENV=DEV
../bin/server -c file://config.toml
