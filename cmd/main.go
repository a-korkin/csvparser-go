package main

import (
	"github.com/a-korkin/csvparser/internals/tools"
)

func main() {
	fileName := "internals/data/ex.csv"
	link := "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	tools.DownloadFile(link, fileName)
	tools.ParseFile(fileName)
}
