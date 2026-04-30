package main

import (
	"context"
	"database/sql"
	_ "embed"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/heroiclabs/nakama-common/runtime"
)

const getGameConfigRPCID = "get_game_config"

//go:embed config/game_config.json
var gameConfigJSON []byte

var cachedGameConfigResponse string

var (
	errGameConfigLoadFailed   = errors.New("unable to load game config")
	errGameConfigInvalidJSON  = errors.New("game config must be valid JSON")
	errGameConfigInvalidValue = errors.New("game config validation failed")
	errGameConfigMarshalFail  = errors.New("unable to marshal game config")
)

type GameConfig struct {
	WelcomeMessage string   `json:"welcomeMessage"`
	XPRate         float64  `json:"xpRate"`
	RarityOptions  []string `json:"rarityOptions"`
}

func getGameConfigRPC(ctx context.Context, logger runtime.Logger, db *sql.DB, nk runtime.NakamaModule, payload string) (string, error) {
	return cachedGameConfigResponse, nil
}

func loadAndCacheGameConfig() error {
	cfg, err := loadGameConfigFromBytes(gameConfigJSON)
	if err != nil {
		return err
	}

	response, err := json.Marshal(cfg)
	if err != nil {
		return fmt.Errorf("%w: %v", errGameConfigMarshalFail, err)
	}

	cachedGameConfigResponse = string(response)
	return nil
}

func loadGameConfig(path string) (GameConfig, error) {
	raw, err := os.ReadFile(path)
	if err != nil {
		return GameConfig{}, fmt.Errorf("%w: %v", errGameConfigLoadFailed, err)
	}

	return loadGameConfigFromBytes(raw)
}

func loadGameConfigFromBytes(raw []byte) (GameConfig, error) {
	var cfg GameConfig
	if err := json.Unmarshal(raw, &cfg); err != nil {
		return GameConfig{}, fmt.Errorf("%w: %v", errGameConfigInvalidJSON, err)
	}

	if err := validateGameConfig(cfg); err != nil {
		return GameConfig{}, fmt.Errorf("%w: %v", errGameConfigInvalidValue, err)
	}

	return cfg, nil
}

func validateGameConfig(cfg GameConfig) error {
	if strings.TrimSpace(cfg.WelcomeMessage) == "" {
		return errors.New("welcomeMessage must be a non-empty string")
	}

	if cfg.XPRate <= 0 {
		return errors.New("xpRate must be greater than 0")
	}

	if len(cfg.RarityOptions) == 0 {
		return errors.New("rarityOptions must not be empty")
	}

	for i, option := range cfg.RarityOptions {
		if strings.TrimSpace(option) == "" {
			return fmt.Errorf("rarityOptions[%d] must be a non-empty string", i)
		}
	}

	return nil
}
