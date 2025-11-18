package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/pelletier/go-toml"
)

type ClientConfig struct {
	MacAddress  string `toml:"macAddress"`
	BroadcastIP string `toml:"broadcastIP"`
}

type SSHConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	User     string `toml:"user"`
	Password string `toml:"password"`
}

type RemoteProgramConfig struct {
	FolderPath  string `toml:"folderPath"`
	ProgramName string `toml:"programName"`
	SudoUser    bool   `toml:"sudoUser"`
}

type Config struct {
	Client        ClientConfig        `toml:"client"`
	SSH           SSHConfig           `toml:"ssh"`
	RemoteProgram RemoteProgramConfig `toml:"remote_program"`
	FileName      string
}

func LoadConfig(filePath string) (Config, error) {
	tree, err := toml.LoadFile(filePath)
	if err != nil {
		return Config{}, fmt.Errorf("failed to load %s: %w", filePath, err)
	}

	var cfg Config
	if err := tree.Unmarshal(&cfg); err != nil {
		return Config{}, fmt.Errorf("failed to parse %s: %w", filePath, err)
	}

	cfg.FileName = filepath.Base(filePath)
	return cfg, nil
}

func LoadConfigs(dir string) ([]Config, error) {
	files, err := os.ReadDir(dir)
	if err != nil {
		return nil, fmt.Errorf("unable to read config directory: %w", err)
	}

	var configs []Config

	for _, f := range files {
		if f.IsDir() || filepath.Ext(f.Name()) != ".toml" {
			continue
		}

		fullPath := filepath.Join(dir, f.Name())

		tree, err := toml.LoadFile(fullPath)
		if err != nil {
			return nil, fmt.Errorf("failed to load %s: %w", fullPath, err)
		}

		var cfg Config
		if err := tree.Unmarshal(&cfg); err != nil {
			return nil, fmt.Errorf("failed to parse %s: %w", fullPath, err)
		}

		cfg.FileName = f.Name()
		configs = append(configs, cfg)
	}

	return configs, nil
}
