# Foreplay

![CI](https://github.com/eberkund/foreplay/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/eberkund/foreplay/branch/master/graph/badge.svg)](https://codecov.io/gh/eberkund/foreplay)
[![Go Report Card](https://goreportcard.com/badge/github.com/eberkund/foreplay)](https://goreportcard.com/report/github.com/eberkund/foreplay)
![GitHub](https://img.shields.io/github/license/eberkund/foreplay)

A precommit hook manager.

### How is this different from precommit?

It makes not attempt to install precommit hooks for you and as a result is a lot simpler. We assume the commands you want to run ahead of each commit have already been setup.

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
