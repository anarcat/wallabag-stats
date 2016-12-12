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
	TableData          []dataRow
	GenTime            time.Time
	GenOverallTime     time.Time
	GenDayTime         time.Time
	GenWeekTime        time.Time
	GenMonthTime       time.Time
	GenYearTime        time.Time
	IsDayGenerated     bool
	IsWeekGenerated    bool
	IsMonthGenerated   bool
	IsYearGenerated    bool
	IsOverallGenerated bool
}

func getLastModTime(filePath string, lastModTime *time.Time) {
	fi, err := os.Stat(filePath)
	if err != nil {
		log.Println("getLastModTime", err)
		return
	}
	*lastModTime = fi.ModTime()
}

func writeTemplateToHTML(wbgStats *WallabagStats, templateName string, isDayGenerated, isWeekGenerated, isMonthGenerated, isYearGenerated, isOverallGenerated bool) {
	if *debug {
		log.Printf("writeTemplateToHTML templateName=%v isDayGenerated=%v", templateName, isDayGenerated)
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

	getLastModTime(chartOverallPath, &td.GenOverallTime)
	getLastModTime(chartDayPath, &td.GenDayTime)
	getLastModTime(chartWeekPath, &td.GenWeekTime)
	getLastModTime(chartMonthPath, &td.GenMonthTime)
	getLastModTime(chartYearPath, &td.GenYearTime)
	td.IsDayGenerated = isDayGenerated
	td.IsWeekGenerated = isWeekGenerated
	td.IsMonthGenerated = isMonthGenerated
	td.IsYearGenerated = isYearGenerated
	td.IsOverallGenerated = isOverallGenerated
	td.GenTime = time.Now()

	htmlSource, err := template.ParseFiles(tmplDirectory+"/"+templateName+".tmpl", tmplDirectory+"/header.tmpl", tmplDirectory+"/footer.tmpl")
	if err != nil {
		log.Println("writeTemplateToHTML", err)
	}
	htmlSource.Execute(f, td)
}

func generateHTML(wbgStats *WallabagStats, isDayGenerated, isWeekGenerated, isMonthGenerated, isYearGenerated, isOverallGenerated bool) {
	if *debug {
		log.Println("generateHTML start")
	}
	err := CopyDir("tmpl/static", outputDirectory)
	if err != nil {
		fmt.Println("error while copying contents from html/ dir to output/ dir. Error:", err)
	}
	writeTemplateToHTML(wbgStats, "data-table", isDayGenerated, isWeekGenerated, isMonthGenerated, isYearGenerated, isOverallGenerated)
	writeTemplateToHTML(wbgStats, "index", isDayGenerated, isWeekGenerated, isMonthGenerated, isYearGenerated, isOverallGenerated)

	if *debug {
		log.Println("generateHTML end")
	}
}
