package utils

import (
	"database/sql"
	// "fmt"
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
	queryArg string
}

func (mdb *MockDB) Exec(query string, args ...interface{}) (sql.Result, error) {
	mdb.queryArg = query
	return nil, nil
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

func TestRunQuery(t *testing.T) {
	want := "SELECT * FROM FOO"
	mockDB := &MockDB{}
	RunQuery(want, mockDB)
	got := mockDB.queryArg

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}

func TestCreateTaskTable(t *testing.T) {
	want := tableCreation
	mockDB := &MockDB{}
	CreateTaskTable(mockDB)
	got := mockDB.queryArg

	if got != want {
		t.Errorf("Expected %q, got %q", want, got)
	}
}
