package main

import (
	"github.com/a-korkin/csvparser/internals/tools"
)

func main() {
	fileName := "ex.csv"
	filePath := tools.PrepareDir(fileName)

	link := "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	tools.DownloadFile(link, filePath)
	tools.ParseFile(filePath)
}
