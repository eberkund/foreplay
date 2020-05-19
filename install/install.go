package install

import (
	"fmt"
	"path"

	"github.com/spf13/afero"
)

var Fs = afero.NewOsFs()

func Init() error {
	contents := `hooks:
#  - id: example sleep
#    run: |
#      sleep 5
`
	return afero.WriteFile(Fs, ".foreplay.yml", []byte(contents), 0755)
}

func Install() error {
	fmt.Println("installing to", path.Join(".git", "hooks"))
	contents := `#!/bin/sh
exec &> /dev/tty
foreplay run
`
	return afero.WriteFile(Fs, PreCommitHookPath(), []byte(contents), 0755)
}

func PreCommitHookPath() string {
	return path.Join(".git", "hooks", "pre-commit")
}
