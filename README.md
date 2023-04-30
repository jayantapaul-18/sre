[![pre-commit](https://img.shields.io/badge/pre--commit-enabled-brightgreen?logo=pre-commit)](https://github.com/pre-commit/pre-commit)
[![go](https://img.shields.io/badge/go-lightblue?logo=go)](https://go.dev/)
[![go](https://img.shields.io/badge/goreleaser-enable-blue?logo=goreleaser)](https://goreleaser.com/)

# SRE CLI tool in Go

Simple go cli tools for sre teams. The main purpose to build this cli to make repetative task simplier. Suppose we want to run healthcheck for many application & if we use UI or go indivisual services then it will be time consueing . So, using cli we can send check command to one serices and that services can all the operation and response back to cli.

-   Command Format :

```bash
sre {command} {subcommand} {args..} {flags..}
```

# Project Development Steps

# Dev setup for goreleaser

```bash
brew install goreleaser
goreleaser init
goreleaser release --snapshot --clean
goreleaser check
goreleaser build --single-target
goreleaser build



```

# Running pre-commit checks

```bash
Install pre-commit and commitizen to use

brew install commitizen
brew install pre-commit
brew install golangci-lint


pre-commit install
pre-commit install --hook-type commit-msg
pre-commit run --all-files

git add .
git status
pre-commit run --all-files
cz c
git commit -m 'feat: health check api with backend api status response'
git push origin main --force
```
