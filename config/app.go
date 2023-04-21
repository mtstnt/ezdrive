package config

import (
	"os"
	"path"
	"path/filepath"
)

// Package config consists of configurations loaded from config file.

var (
	ExecutableDir string
	WorkingDir    string

	CachedDir string
	CredsDir  string
	TokensDir string
)

func LoadAppConfig() error {
	executableDir, err := os.Executable()
	if err != nil {
		return err
	}
	executableDir = path.Dir(executableDir)

	workingDir, err := os.Getwd()
	if err != nil {
		return err
	}

	cachedDir := filepath.Join(executableDir, "cached")
	credsDir := filepath.Join(cachedDir, "creds")
	tokensDir := filepath.Join(cachedDir, "tokens")

	// Create folder if it doesnt exist yet. No exception for uncreated folder.
	for _, thePath := range []string{credsDir, tokensDir} {
		if err := os.MkdirAll(thePath, 0644); err != nil {
			return err
		}
	}

	ExecutableDir = executableDir
	WorkingDir = workingDir

	CachedDir = cachedDir
	CredsDir = credsDir
	TokensDir = tokensDir

	return nil
}
