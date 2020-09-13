package db

import (
	"database/sql"
	"fmt"
	"reflect"
	"testing"
)

type mockFileCreator struct {
	createCalled int
}

func (mfc *mockFileCreator) create(filepath string) {
	mfc.createCalled += 1
}

func mockSQLOpen(driver, conn string) (*sql.DB, error) {
	return nil, nil
}

type MockDB struct {
	sqlExecArg  string
	sqlExecArgs []interface{}

	sqlQueryArg  string
	sqlQueryArgs []interface{}
}

func (mdb *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	mdb.sqlExecArg = query
	mdb.sqlExecArgs = args
	return nil, nil
}

func (mdb *MockDB) Query(query string, args ...interface{}) (rowScanner, error) {
	mdb.sqlQueryArg = query
	mdb.sqlQueryArgs = args
	return &mockRowScanner{rows: 1}, nil
}

func (mdb *MockDB) Close() error {
	return nil
}

func TestCreateSQLiteDB(t *testing.T) {
	mfc := &mockFileCreator{}
	oldSqlOpen := sqlOpen
	defer func() { sqlOpen = oldSqlOpen }()
	sqlOpen = mockSQLOpen

	CreateSQLiteDB("/tmp/foo.db", mfc)

	if mfc.createCalled != 1 {
		t.Errorf("`fileCreator.create` was called %d times, expected %d", mfc.createCalled, 1)
	}

}

func TestSQLExec(t *testing.T) {
	want := "SELECT * FROM FOO"
	mockDB := &MockDB{}
	SQLExec(want, mockDB)
	got := mockDB.sqlExecArg

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestCreateTaskTable(t *testing.T) {
	want := tableCreation
	mockDB := &MockDB{}
	CreateTaskTable(mockDB)
	got := mockDB.sqlExecArg

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

// Implement mock rowScanner
type mockRowScanner struct {
	rowScanner
	rows int
}

func (mrs *mockRowScanner) Next() bool {
	for mrs.rows > 0 {
		mrs.rows -= 1
		return true
	}
	return false
}

func (mrs *mockRowScanner) Close() error {
	return nil
}

type mockData struct {
	taskID int
	title  string
	done   bool
}

func (mrs *mockRowScanner) Scan(args ...interface{}) error {
	taskIDPtr := args[0].(*int)
	*taskIDPtr = 10

	titlePtr := args[1].(*string)
	*titlePtr = "Random task"

	boolPtr := args[2].(*bool)
	*boolPtr = false

	return nil
}

func TestGetTasks(t *testing.T) {
	mockDB := &MockDB{}
	got := GetTasks(mockDB, false)
	want := []TaskStruct{TaskStruct{10, "Random task", false}}

	if !reflect.DeepEqual(got, want) {
		t.Errorf("Got %v, expected %v", got, want)
	}

	if mockDB.sqlQueryArg != SelectStatement {
		t.Errorf("Got %v, expected %v", mockDB.sqlQueryArg, SelectStatement)
	}

}

func TestAddTask(t *testing.T) {
	mockDB := &MockDB{}
	AddTask(mockDB, "Test title", "2020")

	if mockDB.sqlExecArg != InsertStatement {
		t.Errorf("Got %v, expected %v", mockDB.sqlExecArg, InsertStatement)
	}

	fmt.Println(mockDB.sqlExecArgs)

}
