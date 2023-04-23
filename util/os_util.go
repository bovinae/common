package util

import (
	"os/exec"

	"github.com/pkg/errors"
)

func ExecOsCommond(name string, arg ...string) error {
	cmd := exec.Command(name, arg...) // `C:\Program Files\Git\git-bash.exe`, "-c", `rm -rf ./storage`
	if err := cmd.Start(); err != nil {
		return errors.Wrap(err, "ExecOsCommond")
	}
	if err := cmd.Wait(); err != nil {
		return errors.Wrap(err, "ExecOsCommond")
	}
	if !cmd.ProcessState.Success() {
		return errors.New("ExecOsCommond failed")
	}
	return nil
}
