package utils

import (
	"fmt"
	"os"
	"path/filepath"
)

func FileExists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

func FindPackageJSON(startDir string) (string, error) {
	dir := startDir
	for {
		pkgPath := filepath.Join(dir, "package.json")
		if FileExists(pkgPath) {
			return pkgPath, nil
		}
		
		parent := filepath.Dir(dir)
		if parent == dir {
			return "", fmt.Errorf("package.json not found")
		}
		dir = parent
	}
}

func GetCurrentDir() string {
	dir, err := os.Getwd()
	if err != nil {
		return "."
	}
	return dir
}
