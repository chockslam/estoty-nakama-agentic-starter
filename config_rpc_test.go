package main

import (
	"encoding/json"
	"os"
	"path/filepath"
	"testing"
)

func TestValidateGameConfig(t *testing.T) {
	t.Parallel()

	tests := []struct {
		name    string
		cfg     GameConfig
		wantErr bool
	}{
		{
			name: "valid config accepted",
			cfg: GameConfig{
				WelcomeMessage: "Welcome to the game!",
				XPRate:         1.5,
				RarityOptions:  []string{"common", "rare", "epic"},
			},
		},
		{
			name: "missing welcome message rejected",
			cfg: GameConfig{
				WelcomeMessage: "",
				XPRate:         1.5,
				RarityOptions:  []string{"common"},
			},
			wantErr: true,
		},
		{
			name: "invalid xpRate rejected",
			cfg: GameConfig{
				WelcomeMessage: "Welcome to the game!",
				XPRate:         0,
				RarityOptions:  []string{"common"},
			},
			wantErr: true,
		},
		{
			name: "empty rarity options rejected",
			cfg: GameConfig{
				WelcomeMessage: "Welcome to the game!",
				XPRate:         1.5,
				RarityOptions:  nil,
			},
			wantErr: true,
		},
		{
			name: "invalid rarity option rejected",
			cfg: GameConfig{
				WelcomeMessage: "Welcome to the game!",
				XPRate:         1.5,
				RarityOptions:  []string{"common", "   "},
			},
			wantErr: true,
		},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			t.Parallel()

			err := validateGameConfig(tt.cfg)
			if tt.wantErr {
				if err == nil {
					t.Fatal("validateGameConfig() expected error, got nil")
				}
				return
			}

			if err != nil {
				t.Fatalf("validateGameConfig() unexpected error = %v", err)
			}
		})
	}
}

func TestLoadGameConfig(t *testing.T) {
	t.Parallel()

	dir := t.TempDir()

	validPath := filepath.Join(dir, "game_config.json")
	validJSON := []byte(`{"welcomeMessage":"Welcome to the game!","xpRate":1.5,"rarityOptions":["common","rare","epic"]}`)
	if err := os.WriteFile(validPath, validJSON, 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	cfg, err := loadGameConfig(validPath)
	if err != nil {
		t.Fatalf("loadGameConfig() unexpected error = %v", err)
	}

	if cfg.WelcomeMessage != "Welcome to the game!" || cfg.XPRate != 1.5 || len(cfg.RarityOptions) != 3 {
		t.Fatalf("loadGameConfig() = %#v, want valid config", cfg)
	}

	invalidPath := filepath.Join(dir, "invalid.json")
	if err := os.WriteFile(invalidPath, []byte(`{"welcomeMessage":"Welcome to the game!","xpRate":`), 0o600); err != nil {
		t.Fatalf("WriteFile() error = %v", err)
	}

	if _, err := loadGameConfig(invalidPath); err == nil {
		t.Fatal("loadGameConfig() expected error for invalid JSON, got nil")
	}
}

func TestGameConfigMarshalsToValidJSON(t *testing.T) {
	t.Parallel()

	cfg := GameConfig{
		WelcomeMessage: "Welcome to the game!",
		XPRate:         1.5,
		RarityOptions:  []string{"common", "rare", "epic"},
	}

	raw, err := json.Marshal(cfg)
	if err != nil {
		t.Fatalf("json.Marshal() error = %v", err)
	}

	var decoded map[string]any
	if err := json.Unmarshal(raw, &decoded); err != nil {
		t.Fatalf("json.Unmarshal() error = %v", err)
	}

	if decoded["welcomeMessage"] != "Welcome to the game!" {
		t.Fatalf("decoded welcomeMessage = %#v, want %q", decoded["welcomeMessage"], "Welcome to the game!")
	}
	if decoded["xpRate"] != 1.5 {
		t.Fatalf("decoded xpRate = %#v, want 1.5", decoded["xpRate"])
	}
}
