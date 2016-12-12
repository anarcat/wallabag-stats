package main

import (
	"fmt"
	"html/template"
	"log"
	"os"
	"time"
)

const tmplDirectory = "tmpl"

type dataRow struct {
	No      int
	Times   time.Time
	Total   float64
	Unread  float64
	Starred float64
}

type templateData struct {
	TableData []dataRow
	GenTime   time.Time
}

func writeTemplateToHTML(wbgStats *WallabagStats, templateName string) {
	if *debug {
		log.Printf("writeTemplateToHTML templateName=%v", templateName)
	}
	f, err := os.Create(outputDirectory + "/" + templateName + ".html")
	if err != nil {
		log.Println("writeTemplateToHTML", err)
	}
	defer func() {
		if err := f.Close(); err != nil {
			log.Println("writeTemplateToHTML", err)
		}
	}()

	var td templateData
	td.TableData = make([]dataRow, len(wbgStats.Times), len(wbgStats.Times))
	var d dataRow
	for i := 0; i < len(wbgStats.Times); i++ {
		d = dataRow{i + 1, wbgStats.Times[i], wbgStats.Total[i], wbgStats.Unread[i], wbgStats.Starred[i]}
		td.TableData[i] = d
	}

	td.GenTime = time.Now()

	htmlTable, err := template.ParseFiles(tmplDirectory+"/"+templateName+".tmpl", tmplDirectory+"/header.tmpl", tmplDirectory+"/footer.tmpl")
	if err != nil {
		log.Println("writeTemplateToHTML", err)
	}
	htmlTable.Execute(f, td)
}

func generateHTML(wbgStats *WallabagStats) {
	if *debug {
		log.Println("generateHTML start")
	}
	err := CopyDir("tmpl/static", outputDirectory)
	if err != nil {
		fmt.Println("error while copying contents from html/ dir to output/ dir. Error:", err)
	}
	writeTemplateToHTML(wbgStats, "data-table")
	writeTemplateToHTML(wbgStats, "index")

	if *debug {
		log.Println("generateHTML end")
	}
}
