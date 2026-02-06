package command

import (
	"fmt"
	"os/exec"
	"strings"

	"go.mattglei.ch/timber"
)

func Run(binary string, args ...string) error {
	out, err := exec.Command(binary, args...).CombinedOutput()
	if err != nil {
		timber.Debug(string(out))
		return fmt.Errorf(
			`running command "%s %s": %w`,
			binary,
			strings.Join(args, " "),
			err,
		)
	}
	return nil
}
