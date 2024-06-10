#!/usr/bin/env bash

mkdir /etc/shell_vault

/opt/shell_vault/bin/client -a "$SHELL_VAULT_HOST" get-key > /etc/shell_vault/user_ca.key
cat /etc/shell_vault/user_ca.key

echo 'TrustedUserCAKeys /etc/shell_vault/user_ca.key' > /etc/ssh/sshd_config.d/
