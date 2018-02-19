package print

import (
	"os"

	"github.com/olekukonko/tablewriter"
)

type tableWriter struct {
	headers []string
	rows    map[string][]string
}

// New returns printer
func New(headers []string, rows map[string][]string) Printer {
	return tableWriter{
		headers: headers,
		rows:    rows,
	}
}

func (t tableWriter) Print() {
	table := tablewriter.NewWriter(os.Stdout)
	table.SetHeader(t.headers)
	// table.AppendBulk(t.rows)

	table.Render()
}
