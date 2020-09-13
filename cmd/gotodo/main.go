package main

import (
	"fmt"
	"github.com/nakhan98/gotodo/db"
	"github.com/nakhan98/gotodo/ui"
	"github.com/nakhan98/gotodo/utils"
	"os"
	// "os/user"
)

const DBFile string = ".gotodo.sqlite3"

func main() {
	checkDB()
	args := os.Args

	dbConn := openDB()
	defer dbConn.Close()

	if len(args) == 1 {
		// Just list tasks if no args provided
		listTasks(dbConn)
	} else if len(args) == 2 && args[1] == "add" {
		addTask(dbConn)
	}

	// if utils.FileExists(dbFilePath) {
	// dbConn = db.DBOpen(dbFilePath)
	// taskList := db.GetTasks(dbConn, true)
	// if len(taskList) == 0 {
	// 	fmt.Println("GoToDo: No tasks found")
	// } else {
	// 	ui.PrintTasks(db.GetTasks(dbConn, true))
	// }

}

func listTasks(dbConn db.SQLDB) {
	taskList := db.GetTasks(dbConn, false)
	if len(taskList) == 0 {
		fmt.Println("GoToDo: No tasks found")
	} else {
		ui.PrintTasks(taskList)
	}
}

func addTask(dbConn db.SQLDB) {
	input := utils.GetTaskTitle()
	datetimeStr := utils.GetTime()
	db.AddTask(dbConn, input, datetimeStr)
	fmt.Println("Added task!")
}

func openDB() db.SQLDB {
	curUser := utils.GetCurUser()
	dbFilePath := utils.ConstructDBFilePath(curUser, DBFile)
	return db.DBOpen(dbFilePath)
}

// Check if DB exists
// If not create it
func checkDB() {
	curUser := utils.GetCurUser()
	dbFilePath := utils.ConstructDBFilePath(curUser, DBFile)

	if utils.FileExists(dbFilePath) {
		return
	}

	fmt.Println("No sqlite db exists!")
	dbConn := db.CreateSQLiteDB(dbFilePath, db.FileCreatorHelper{})
	db.CreateTaskTable(dbConn)
	defer dbConn.Close()
}
