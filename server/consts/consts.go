package consts

import _ "embed"

//go:embed config.toml
var ConfigTemplate string

//go:embed server.env
var ServerEnv string

const UNIXEtcFolder = "/etc/shell_vault"
const UNIXVarFolder = "/var/lib/shell_vault"
const UNIXVarLogFolder = "/var/log/shell_vault"

const WindowsEtcFolder = "C:\\ProgramData\\ShellVault"
const WindowsVarFolder = "C:\\ProgramData\\ShellVault"
const WindowsVarLogFolder = "C:\\ProgramData\\ShellVault\\logs"

const DefaultConfigFile = "config.toml"
const DefaultServerEnvFile = "server.env"
