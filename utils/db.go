package utils

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/mattn/go-sqlite3"
)

var FileCreate = os.Create
var sqlOpen = sql.Open

type SQLDB interface {
	Exec(query string, args ...interface{}) (sql.Result, error)
}

const tableCreation = `CREATE TABLE Task (
	"taskID" INTEGER NOT NULL PRIMARY KEY AUTOINCREMENT,
	"title" TEXT,
	"datetime" DATETIME,
	"done" BOOLEAN)
	`

type fileCreator interface {
	create(filepath string)
}

type FileCreatorHelper struct{}

func (fc FileCreatorHelper) create(filepath string) {
	// log.Println("Creating sqlite-database.db...")
	fh, err := FileCreate(filepath) // Create SQLite file
	if err != nil {
		panic(err)
	}
	fh.Close()
	// log.Println("sqlite-database.db created")
}

func RunQuery(sqlStatement string, db SQLDB) error {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

func CreateSQLiteDB(filepath string, fc fileCreator) *sql.DB {
	fmt.Println("Creating sql db at:", filepath)
	fc.create(filepath)
	db, err := sqlOpen("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	return db
}

func CreateTaskTable(db SQLDB) {
	RunQuery(tableCreation, db)
}

func GetTasks(db SQLDB, completed bool) {
	//
}
