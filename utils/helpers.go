package utils

import (
	"bufio"
	"fmt"
	"os"
	"os/user"
	"path/filepath"
)

func GetTaskTitle() string {
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Print("Enter task: ")
	scanner.Scan()
	return scanner.Text()
}

func GetCurUser() *user.User {
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	return curUser
}

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
