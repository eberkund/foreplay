# Foreplay

A precommit hook manager.

### How is this different from precommit?

It makes not attempt to install precommit hooks for you and as a result is a lot simpler. We assume the commands you want to run ahead of each commit have already been setup.

Config syntax inspired by VS Code.

```yaml
hooks:
  - name: eslint
    run: npm lint
    working-directory: ./frontend
    depends-on: golang-ci-lint
    shell: node

  - id: golangci-lint
    run: golangci-lint run
    working-directory: ./
    shell: powershell

  - name: test
    script: scripts/bash-example.sh
    working-directory: ./frontend
    shell: /bin/bash
```

# Roadmap

## v0.1
- Installs a shim in .git/hooks
- Loads hooks from yml file

## v0.2
- Support alternative shells to invoke scripts

## v0.3
- Support embedded Go interpretter