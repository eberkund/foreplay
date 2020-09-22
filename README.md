# Foreplay

![CI](https://github.com/eberkund/foreplay/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/eberkund/foreplay/branch/master/graph/badge.svg)](https://codecov.io/gh/eberkund/foreplay)
[![Go Report Card](https://goreportcard.com/badge/github.com/eberkund/foreplay)](https://goreportcard.com/report/github.com/eberkund/foreplay)
![GitHub](https://img.shields.io/github/license/eberkund/foreplay)
[![Snap Status](https://snapcraft.io/foreplay/badge.svg)](https://snapcraft.io/foreplay)

A pre-commit hook manager.

### How is this different from pre-commit?

It makes not attempt to install pre-commit hooks for you and as a result is a lot simpler. We assume the commands you want to run ahead of each commit have already been setup.

## Two output methods

For when looks are more important than personality.

```yaml
spinners: true
```

![GitHub](./example.svg)

Github Actions inspired configuration syntax.

```yaml
hooks:
  - name: eslint
    run: |
      cd frontend
      npm lint

  - id: golangci-lint
    run: golangci-lint run
```

## Installation

### macOS

```shell script
brew tap eberkund/foreplay
brew install foreplay
```

### Linux

```shell script
snap install foreplay
```
Or download install the _.deb_ from the releases page.

### Windows

```shell script
scoop bucket add app https://github.com/eberkund/scoop-foreplay.git
scoop install foreplay
```

### Go

```shell script
go install -u github.com/eberkund/foreplay
```
