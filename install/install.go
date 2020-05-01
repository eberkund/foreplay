package install

import (
	"fmt"
	"os"
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
	preCommitHookPath, err := PreCommitHookPath()
	if err != nil {
		return err
	}
	contents := `#!/usr/bin/env bash
exec foreplay run
`
	return afero.WriteFile(fs, preCommitHookPath, []byte(contents), 0755)
}

func PreCommitHookPath() (string, error) {
	wd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return path.Join(wd, ".git", "hooks", "pre-commit"), nil
}
