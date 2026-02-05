package conf

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml/v2"
	"go.mattglei.ch/timber"
)

func Read() (Config, error) {
	home, err := os.UserHomeDir()
	if err != nil {
		return Config{}, fmt.Errorf("loading user's home directory: %w", err)
	}
	path := filepath.Join(home, ".config", "notes", "config.toml")

	bin, err := os.ReadFile(path)
	if err != nil {
		return Config{}, fmt.Errorf("reading config file at %s: %w", path, err)
	}

	var config Config
	err = toml.Unmarshal(bin, &config)
	if err != nil {
		timber.Debug(string(bin))
		return Config{}, fmt.Errorf("parsing TOML file at %s: %w", path, err)
	}
	return config, nil
}
