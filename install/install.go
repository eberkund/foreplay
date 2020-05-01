package install

import (
	"fmt"
	"path"

	"github.com/spf13/afero"
)

var fs = afero.NewOsFs()

func Init() error {
	contents := `hooks:
#  - id: golangci-lint
#    run: run
`
	return afero.WriteFile(fs, ".foreplay.yml", []byte(contents), 0755)
}

func Install() error {
	fmt.Println("installing to", path.Join(".git", "hooks"))
	contents := `#!/usr/bin/env bash
exec foreplay run
`
	return afero.WriteFile(fs, PreCommitHookPath(), []byte(contents), 0755)
}

func PreCommitHookPath() string {
	return path.Join(".git", "hooks", "pre-commit")
}
