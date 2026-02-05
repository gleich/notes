package conf

import (
	"fmt"
	"os"
)

type Config struct {
	Path string
}

func (c Config) GoToPath() error {
	err := os.Chdir(c.Path)
	if err != nil {
		return fmt.Errorf("changing directory to %s: %w", c.Path, err)
	}
	return nil
}
