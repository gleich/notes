package conf

import (
	"fmt"
	"os"
	"path/filepath"
)

type Config struct {
	Path string
}

func (c Config) ValidatePath() error {
	pwd, err := os.Getwd()
	if err != nil {
		return fmt.Errorf("getting working working directory: %w", err)
	}
	if filepath.Clean(pwd) != filepath.Clean(c.Path) {
		return fmt.Errorf("please run program from %s", c.Path)
	}
	return nil
}
