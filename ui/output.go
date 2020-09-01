package ui

import (
	// "fmt"
	"github.com/nakhan98/gotodo/db"
	"github.com/olekukonko/tablewriter"
	"os"
	"strconv"
)

type TablerWriterWrapper interface {
	setHeader([]string)
	renderTable()
	addRow([]string)
}

type TableWriter struct {
	tw *tablewriter.Table
}

func (t *TableWriter) setHeader(headers []string) {
	t.tw.SetHeader([]string{"TaskID", "Task", "Done"})
}

func (t *TableWriter) renderTable() {
	t.tw.Render()
}

func (t *TableWriter) addRow(row []string) {
	t.tw.Append(row)
}

func PrintTasks(tasks []db.TaskStruct) {
	tableWriter := TableWriter{tw: tablewriter.NewWriter(os.Stdout)}
	printTasks(tasks, &tableWriter)
}

func printTasks(tasks []db.TaskStruct, tableWriter TablerWriterWrapper) {
	// table := Table{tw: tablewriter.NewWriter(os.Stdout)}
	tableWriter.setHeader([]string{"TaskID", "Task", "Done"})

	for _, t := range tasks {
		task := []string{strconv.Itoa(t.TaskID), t.Title, strconv.FormatBool(t.Done)}
		tableWriter.addRow(task)
	}
	tableWriter.renderTable()

}
