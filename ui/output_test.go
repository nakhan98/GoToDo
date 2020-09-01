package ui

import (
	"github.com/nakhan98/gotodo/db"
	"reflect"
	"testing"
)

type mockTable struct {
	TablerWriterWrapper
	header       []string
	rows         [][]string
	renderCalled bool
}

func (mt *mockTable) setHeader(header []string) {
	mt.header = header
}

func (mt *mockTable) addRow(row []string) {
	mt.rows = append(mt.rows, row)
}

func (mt *mockTable) renderTable() {
	mt.renderCalled = true
}

func TestPrintTasks(t *testing.T) {
	testTasks := []db.TaskStruct{db.TaskStruct{10, "Do this", false}}
	headersExpected := []string{"TaskID", "Task", "Done"}
	rowExpected := []string{"10", "Do this", "false"}

	mockTableWriter := &mockTable{}
	printTasks(testTasks, mockTableWriter)

	if !reflect.DeepEqual(mockTableWriter.header, headersExpected) {
		t.Errorf("Expected %q, got %q", headersExpected, mockTableWriter.header)
	}

	if !reflect.DeepEqual(mockTableWriter.rows[0], rowExpected) {
		t.Errorf("Expected %q, got %q", rowExpected, mockTableWriter.rows)
	}

	if !mockTableWriter.renderCalled {
		t.Errorf("Expected %v, got %v", true, mockTableWriter.renderCalled)
	}

}
