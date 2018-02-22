package print

import (
	"fmt"
	"os"

	"github.com/olekukonko/tablewriter"
)

type tableWriter struct {
	headers []string
	rows    []map[string]interface{}
}

// New returns printer
func New(headers []string, rows []map[string]interface{}) Printer {
	return tableWriter{
		headers: headers,
		rows:    rows,
	}
}

func (t tableWriter) Print() {
	var rows [][]string
	for _, r := range t.rows {
		fmt.Printf("%+v", r)
		// t := t.ttype
		// rows = append(rows, r)
	}
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.headers)
	table.AppendBulk(rows)

	table.Render()
}
