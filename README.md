# Dotenv - Run a command with env vars injected from one or more .env files

## Installation

```bash
go install github.com/pomdtr/dotenv
```

## Usage

```text
Run a command with env vars injected from one or more .env files

Usage:
  dotenv [flags] command [args...]

Examples:
dotenv -- YOUR_COMMAND --YOUR-FLAG
doten run --command "YOUR_COMMAND && YOUR_OTHER_COMMAND"

Flags:
  -c, --command string      command to run
      --completion string   generate completion script
  -e, --env strings         env files to load
  -h, --help                help for dotenv
      --preserve-env        preserve existing environment variables
```

> **Note**
> If you pass the `--command` flag, the command will be executed in a subshell. The shell will be determined by the `SHELL` environment variable. If it is not set, `sh` will be used as a fallback. If you want to use a different shell, try: `dotenv -- zsh -c "YOUR_COMMAND"`.

## Shell Completions

### Bash

```bash
dotenv --completion bash > /etc/bash_completion.d/dotenv # Linux
dotenv --completion bash > $(brew --prefix)/etc/bash_completion.d/dotenv # macOS
```

### Zsh

```zsh
dotenv completion zsh > "${fpath[1]}/_dotenv" # Linux
dotenv completion zsh > $(brew --prefix)/share/zsh/site-functions/_dotenv # macOS
```

### Fish

```fish
dotenv --completion fish > ~/.config/fish/completions/dotenv.fish
```

### Powershell

```powershell
dotenv --completion powershell | Out-String | Invoke-Expression
```
