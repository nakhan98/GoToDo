package main

import (
	"fmt"
	"github.com/nakhan98/gotodo/utils"
	"os/user"
)

const DBFile string = ".gotodo.sqlite3"

func main() {
	fmt.Println("GoToDo")
	curUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	dbFilePath := utils.ConstructDBFilePath(curUser, DBFile)
	if utils.FileExists(dbFilePath) {
		fmt.Println("DB File exists")
	} else {
		fmt.Println("DB File does not exist")
		db := utils.CreateSQLiteDB(dbFilePath, utils.FileCreatorHelper{})
		utils.CreateTaskTable(db)

	}
}
