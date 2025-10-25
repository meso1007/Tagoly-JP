package config

import (
	"encoding/json"
	"os"
	"path/filepath"
)

type Config struct {
	CustomTags []string `json:"customTags"`
}

func LoadConfig() (*Config, error) {
	cfg := &Config{}

	// プロジェクト直下
	if _, err := os.Stat(".tagolyrc"); err == nil {
		data, _ := os.ReadFile(".tagolyrc")
		json.Unmarshal(data, cfg)
		return cfg, nil
	}

	// ホームディレクトリ直下
	home, err := os.UserHomeDir()
	if err != nil {
		return cfg, err
	}
	file := filepath.Join(home, ".tagolyrc")
	if _, err := os.Stat(file); os.IsNotExist(err) {
		return cfg, nil
	}
	data, err := os.ReadFile(file)
	if err != nil {
		return cfg, err
	}
	json.Unmarshal(data, cfg)
	return cfg, nil
}
