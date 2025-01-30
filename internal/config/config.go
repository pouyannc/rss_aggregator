package config

import (
	"encoding/json"
	"os"
)

const configFileName = ".gatorconfig.json"

type Config struct {
	DbURL           string `json:"db_url"`
	CurrentUsername string `json:"current_user_name"`
	RSSURL          string
}

func (cfg *Config) SetUser(username string) error {
	cfg.CurrentUsername = username

	dat, err := json.Marshal(cfg)
	if err != nil {
		return err
	}

	homeDir, err := os.UserHomeDir()
	if err != nil {
		return err
	}

	cfgPath := homeDir + "/" + configFileName

	err = os.WriteFile(cfgPath, dat, 0777)
	if err != nil {
		return err
	}

	return nil
}

func Read() (Config, error) {
	homeDir, err := os.UserHomeDir()
	if err != nil {
		return Config{}, err
	}

	cfgPath := homeDir + "/" + configFileName

	dat, err := os.ReadFile(cfgPath)
	if err != nil {
		return Config{}, err
	}

	var cfg Config
	err = json.Unmarshal(dat, &cfg)
	if err != nil {
		return Config{}, err
	}

	return cfg, nil
}
