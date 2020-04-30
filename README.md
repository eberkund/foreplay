# Foreplay

![CI](https://github.com/eberkund/foreplay/workflows/CI/badge.svg)
[![codecov](https://codecov.io/gh/eberkund/foreplay/branch/master/graph/badge.svg)](https://codecov.io/gh/eberkund/foreplay)
[![Go Report Card](https://goreportcard.com/badge/github.com/eberkund/foreplay)](https://goreportcard.com/report/github.com/eberkund/foreplay)

A precommit hook manager.

### How is this different from precommit?

It makes not attempt to install precommit hooks for you and as a result is a lot simpler. We assume the commands you want to run ahead of each commit have already been setup.

## Two output methods

For when looks are more important than personality.

```yaml
spinners: true
```

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

```
# Register hooks
foreplay install

# Manual run
foreplay run

# Example output
$ foreplay run
+---------------+---+
| golangci-lint | ✓ |
| goreleaser    | ✓ |
| npm test      | ✗ |
+---------------+---+
```

# Roadmap

## v0.1
- Installs a shim in .git/hooks
- Loads hooks from yml file

## v0.2
- Support alternative shells to invoke scripts

## v0.3
- Support embedded Go interpretter
