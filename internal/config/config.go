package config

import (
	"encoding/json"
	"fmt"
	"os"
)

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	cfgPath := homeDir + "/.gatorconfig.json"

	dat, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return Config{}, err
	}

	fmt.Println(dat)
	return cfg, nil
}
