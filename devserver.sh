#!/usr/bin/env bash

make server

# first setup
#../bin/server -first-setup

export SHELL_VAULT_ENV=DEV
export SHELL_VAULT_KEK='testkekkey'
export SHELL_VAULT_ROOT_PASSWORD='testtest'
bin/server -run -config dev_root/etc/shell_vault -override-root ./dev_root
