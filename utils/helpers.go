package utils

import (
	// "fmt"
	"os"
	"os/user"
	"path/filepath"
)

func ConstructDBFilePath(u *user.User, dbFileName string) string {
	userDir := u.HomeDir
	return filepath.Join(userDir, dbFileName)
}

func FileExists(filepath string) bool {
	if _, err := os.Stat(filepath); err == nil {
		return true
	}
	return false

}
