package config

import (
	"os"
	"path/filepath"
	"testing"
)

func TestLoad(t *testing.T) {
	// Create a temporary config file
	tmpDir := t.TempDir()
	configPath := filepath.Join(tmpDir, "config.yml")

	configContent := `customCommands:
  - key: "D"
    description: "Test command"
    context: "list"
    command: "echo hello"
  - key: "C"
    description: "Another command"
    context: "detail"
    command: "echo {{.ID}}"
`
	if err := os.WriteFile(configPath, []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	// Override config path for testing
	originalUserConfigDir := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	defer os.Setenv("XDG_CONFIG_HOME", originalUserConfigDir)

	// Create the lazybeads subdirectory
	if err := os.MkdirAll(filepath.Join(tmpDir, "lazybeads"), 0755); err != nil {
		t.Fatalf("failed to create config dir: %v", err)
	}
	if err := os.WriteFile(filepath.Join(tmpDir, "lazybeads", "config.yml"), []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	cfg, err := Load()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if len(cfg.CustomCommands) != 2 {
		t.Errorf("expected 2 custom commands, got %d", len(cfg.CustomCommands))
	}

	if cfg.CustomCommands[0].Key != "D" {
		t.Errorf("expected first command key to be 'D', got '%s'", cfg.CustomCommands[0].Key)
	}

	if cfg.CustomCommands[0].Context != "list" {
		t.Errorf("expected first command context to be 'list', got '%s'", cfg.CustomCommands[0].Context)
	}

	if cfg.CustomCommands[1].Context != "detail" {
		t.Errorf("expected second command context to be 'detail', got '%s'", cfg.CustomCommands[1].Context)
	}
}

func TestLoadNoConfig(t *testing.T) {
	// Point to a nonexistent config directory
	tmpDir := t.TempDir()
	originalUserConfigDir := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	defer os.Setenv("XDG_CONFIG_HOME", originalUserConfigDir)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("expected no error for missing config, got: %v", err)
	}

	if len(cfg.CustomCommands) != 0 {
		t.Errorf("expected empty custom commands for missing config, got %d", len(cfg.CustomCommands))
	}
}

func TestDefaultContext(t *testing.T) {
	tmpDir := t.TempDir()
	if err := os.MkdirAll(filepath.Join(tmpDir, "lazybeads"), 0755); err != nil {
		t.Fatalf("failed to create config dir: %v", err)
	}

	// Config with no context specified
	configContent := `customCommands:
  - key: "X"
    description: "No context"
    command: "echo test"
`
	if err := os.WriteFile(filepath.Join(tmpDir, "lazybeads", "config.yml"), []byte(configContent), 0644); err != nil {
		t.Fatalf("failed to write test config: %v", err)
	}

	originalUserConfigDir := os.Getenv("XDG_CONFIG_HOME")
	os.Setenv("XDG_CONFIG_HOME", tmpDir)
	defer os.Setenv("XDG_CONFIG_HOME", originalUserConfigDir)

	cfg, err := Load()
	if err != nil {
		t.Fatalf("failed to load config: %v", err)
	}

	if len(cfg.CustomCommands) != 1 {
		t.Fatalf("expected 1 custom command, got %d", len(cfg.CustomCommands))
	}

	// Should default to "list"
	if cfg.CustomCommands[0].Context != "list" {
		t.Errorf("expected default context to be 'list', got '%s'", cfg.CustomCommands[0].Context)
	}
}
