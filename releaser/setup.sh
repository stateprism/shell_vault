#!/usr/bin/env bash

mkdir -p /opt/shell_vault/bin

function getPlatform() {
    # Get the platform of the current machine
    case "$(uname -m)" in
        "x86_64")
            platform="amd64"
            ;;
        "aarch64")
            platform="arm64"
            ;;
        *)
            echo "Unsupported platform"
            exit 1
            ;;
    esac
    # get the OS of the current machine

}

curl -SsL https://dist.stateprism.com/shell-vault/latest/linux_amd64.tar.gz | tar -xz -C /opt

tee -a /etc/systemd/system/shell_vault.service <<EOF
[Unit]
Description=Shell Vault SSH Certificate Authority
After=network.target

[Service]
Type=simple
ExecStart=/opt/shell-vault/bin/server
ConfigurationDirectory=/etc/shell-vault
EnvironmentFile=/etc/shell-vault/server.env
Restart=always

[Install]
Alias=shell-vault.service

EOF

systemctl daemon-reload
systemctl enable --now shell-vault
systemctl status shell-vault
