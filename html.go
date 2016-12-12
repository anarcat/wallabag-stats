package main

import (
	"html/template"
	"log"
	"os"
	"time"
)

type dataRow struct {
	No      int
	Times   time.Time
	Total   float64
	Unread  float64
	Starred float64
}

func writeDataTableHTML(wbgStats *WallabagStats) {
	fo, err := os.Create(outputDirectory + "/data-table.html")
	if err != nil {
		log.Println(err)
	}
	defer func() {
		if err := fo.Close(); err != nil {
			log.Println(err)
		}
	}()

	tableData := make([]dataRow, len(wbgStats.Times), len(wbgStats.Times))
	var d dataRow
	for i := 0; i < len(wbgStats.Times); i++ {
		d = dataRow{i + 1, wbgStats.Times[i], wbgStats.Total[i], wbgStats.Unread[i], wbgStats.Starred[i]}
		tableData[i] = d
	}
	const tableTpl = `<table>
		<th>
			<td>No.</td>
			<td>Date</td>
			<td>Total</td>
			<td>Unread</td>
			<td>Starred</td>
		</th>
		{{range .}}<tr><td>{{ .No }}</td><td>{{ .Times }}</td><td>{{ .Total }}</td><td>{{ .Unread }}</td><td>{{ .Starred }}</td></tr>{{else}}<div><strong>no data</strong></div>{{end}}
	</table>`
	htmlTable, err := template.New("data-table").Parse(tableTpl)
	if err != nil {
		log.Println("writeDataTableHTML error", err)
	}
	htmlTable.Execute(fo, tableData)
}

func generateHTML(wbgStats *WallabagStats) {
	if *debug {
		log.Println("generateHTML start")
	}
	writeDataTableHTML(wbgStats)

	if *debug {
		log.Println("generateHTML end")
	}
}
