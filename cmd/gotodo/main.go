package main

import (
	"fmt"
	"github.com/nakhan98/gotodo/db"
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
	var dbConn db.SQLDB
	dbFilePath := utils.ConstructDBFilePath(curUser, DBFile)
	if utils.FileExists(dbFilePath) {
		fmt.Println("DB File exists")
		dbConn = db.DBOpen(dbFilePath)
	} else {
		fmt.Println("DB File does not exist")
		sqliteDB := db.CreateSQLiteDB(dbFilePath, db.FileCreatorHelper{})
		db.CreateTaskTable(sqliteDB)
		dbConn = sqliteDB
	}
	db.GetTasks(dbConn, true)
}
