package utils

import (
	"os/user"
	"runtime"
	"testing"
)

func mockHomeDir() string {
	return "/home/tester/"
}

func TestConstructDBFilePath(t *testing.T) {
	want := "/home/tester/DBFILE.sqlite3"

	u, _ := user.Current()
	orgHome := u.HomeDir
	u.HomeDir = "/home/tester/"

	defer func() {
		u.HomeDir = orgHome
	}()

	got := ConstructDBFilePath(u, "DBFILE.sqlite3")

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func getTestFilePath() string {
	_, filename, _, _ := runtime.Caller(0)
	return filename
}

func TestFileExistsOK(t *testing.T) {
	fp := getTestFilePath()

	want := true
	got := FileExists(fp)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}

func TestFileExistsFail(t *testing.T) {
	fp := getTestFilePath() + "_non_existent"

	want := false
	got := FileExists(fp)

	if got != want {
		t.Errorf("got %v, want %v", got, want)
	}

}
