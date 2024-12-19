package main

import (
	"github.com/a-korkin/csvparser/internals/tools"
	"path/filepath"
	"sync"
)

func worker(filePath, fileName, link string) {
	tools.DownloadFile(link, filePath)
	tools.ParseFile(fileName, filePath)
}

func main() {
	files := make(map[string]string, 2)
	files["survey.csv"] = "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	files["finance.csv"] = "https://www.stats.govt.nz/assets/Uploads/Business-operations-survey/Business-operations-survey-2022/Download-data/business-operations-survey-2022-business-finance.csv"

	var wg sync.WaitGroup
	dirName := tools.PrepareDir()
	for fileName, link := range files {
		wg.Add(1)
		go func() {
			defer wg.Done()
			filePath := filepath.Join(dirName, fileName)
			worker(filePath, fileName, link)
		}()
	}
	wg.Wait()
}
