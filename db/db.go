package db

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
	Query(query string, args ...interface{}) (rowScanner, error)
	Close() error
}

type rowScanner interface {
	Next() bool
	Close() error
	Scan(...interface{}) error
}

type sqlWrapper struct {
	db *sql.DB
	SQLDB
}

func (sw *sqlWrapper) Query(query string, args ...interface{}) (rowScanner, error) {
	return sw.db.Query(query, args...)
}

// type scanner interface {
// 	Scan(dest ...interface{}) error
// }

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

func SQLExec(sqlStatement string, db SQLDB) error {
	_, err := db.Exec(sqlStatement)
	if err != nil {
		return err
	}
	return nil
}

func DBOpen(filepath string) SQLDB {
	db, err := sqlOpen("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	return &sqlWrapper{db: db}
}

func CreateSQLiteDB(filepath string, fc fileCreator) SQLDB {
	fmt.Println("Creating sql db at:", filepath)
	fc.create(filepath)
	db, err := sqlOpen("sqlite3", filepath)
	if err != nil {
		panic(err)
	}
	return &sqlWrapper{db: db}
}

func CreateTaskTable(db SQLDB) {
	SQLExec(tableCreation, db)
}

type TaskStruct struct {
	taskID int
	title  string
	done   bool
}

const SelectStatement = "SELECT taskID, title, done FROM Task WHERE done = ?"

// Use struct for return val instead
func GetTasks(db SQLDB, done bool) []TaskStruct {
	var tasks []TaskStruct
	rows, err := db.Query(SelectStatement, done)

	if err != nil {
		panic(err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			taskID int
			title  string
			done   bool
		)
		if err := rows.Scan(&taskID, &title, &done); err != nil {
			panic(err)
		}
		fmt.Printf("taskID %d, title is %q, done %v\n", taskID, title, done)
		result := TaskStruct{taskID: taskID, title: title, done: done}
		tasks = append(tasks, result)
	}

	return tasks
}
