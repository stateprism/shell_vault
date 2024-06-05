#!/usr/bin/env bash

make server

cd dev_root || exit 1

export SHELL_VAULT_ENV=DEV
../bin/server -c etc/prisma_ca
