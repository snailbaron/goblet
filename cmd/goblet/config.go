package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
)

func ConfigPath() string {
	configHome := os.Getenv("XDG_CONFIG_HOME")
	if configHome == "" {
		home, err := os.UserHomeDir()
		if err != nil {
			panic(fmt.Sprintf("cannot find user home directory: %v", err))
		}
		configHome = filepath.Join(home, ".config")
	}

	return filepath.Join(configHome, "goblet", "config.json")
}

type Config struct {
	FPS             int      `json:"fps"`
	WindowWidth     int      `json:"window_width"`
	WindowHeight    int      `json:"window_height"`
	BackgroundColor [4]uint8 `json:"background_color"`
}

//nolint:mnd
var defaultConfig = Config{
	FPS:             60,
	WindowWidth:     800,
	WindowHeight:    600,
	BackgroundColor: [4]uint8{30, 30, 30, 255},
}

func LoadConfig() (*Config, error) {
	return LoadConfigFromFile(ConfigPath())
}

func LoadConfigFromFile(path string) (*Config, error) {
	bs, err := os.ReadFile(filepath.Clean(path))
	if err != nil {
		if errors.Is(err, fs.ErrNotExist) {
			bs = []byte("{}")
		} else {
			return nil, err
		}
	}

	c := defaultConfig
	if err := json.Unmarshal(bs, &c); err != nil {
		return nil, err
	}

	return &c, nil
}

func (c *Config) Save() error {
	return c.SaveToFile(ConfigPath())
}

func (c *Config) SaveToFile(path string) error {
	bs, err := json.MarshalIndent(c, "", "    ")
	if err != nil {
		return err
	}

	return writeFileAtomically(bs, path)
}
