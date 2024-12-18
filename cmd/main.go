package main

import (
	"github.com/a-korkin/csvparser/internals/tools"
	"log"
	"os"
	"path/filepath"
)

func main() {
	dirName := "data"
	fileName := "ex.csv"
	err := os.MkdirAll(dirName, os.ModePerm)
	if err != nil {
		log.Fatalf("failed to create dir: %s", err)
	}

	filePath := filepath.Join(dirName, fileName)

	link := "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	tools.DownloadFile(link, filePath)
	tools.ParseFile(filePath)
}
