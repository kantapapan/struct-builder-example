package configs

import (
	"errors"
	"os"

	"github.com/BurntSushi/toml"
)

// ConfigPath Env Path
const ConfigPath = "DEMO_CONFIG_PATH"

// Load は、設定ファイルをロードしConfigを返却します
func Load() (c Config, e error) {
	configPath := os.Getenv(ConfigPath)
	_, err := toml.DecodeFile(configPath, &c)
	if err != nil {
		return c, errors.New("The configuration file was not found")
	}
	return c, nil
}

// Config ...
type Config struct {
	Storage StorageConfig `toml:"storage"`
}

// StorageConfig ...
type StorageConfig struct {
	Path string `toml:"pass"`
}
