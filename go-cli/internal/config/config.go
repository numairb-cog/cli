package config

import (
	"os"
	"path/filepath"
)

type Config struct {
	Registry      string
	Global        bool
	SaveDev       bool
	SaveExact     bool
	IgnoreScripts bool
	LogLevel      string
	Force         bool
	DryRun        bool
	Workspaces    []string
}

func NewConfig() *Config {
	return &Config{
		Registry:      "https://registry.npmjs.org",
		Global:        false,
		SaveDev:       false,
		SaveExact:     false,
		IgnoreScripts: false,
		LogLevel:      "notice",
		Force:         false,
		DryRun:        false,
		Workspaces:    []string{},
	}
}

func GetNpmRoot() string {
	if global := os.Getenv("NPM_CONFIG_PREFIX"); global != "" {
		return global
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".npm")
}

func GetGlobalPrefix() string {
	if prefix := os.Getenv("NPM_CONFIG_PREFIX"); prefix != "" {
		return prefix
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".npm-global")
}

func GetCacheDir() string {
	if cache := os.Getenv("NPM_CONFIG_CACHE"); cache != "" {
		return cache
	}
	home, _ := os.UserHomeDir()
	return filepath.Join(home, ".npm")
}
