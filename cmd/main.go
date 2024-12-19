package main

import (
	"github.com/a-korkin/csvparser/internals/tools"
)

func main() {
	files := make(map[string]string, 2)
	files["survey.csv"] = "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	files["finance.csv"] = "https://www.stats.govt.nz/assets/Uploads/Business-operations-survey/Business-operations-survey-2022/Download-data/business-operations-survey-2022-business-finance.csv"
	// link := "https://www.stats.govt.nz/assets/Uploads/Annual-enterprise-survey/Annual-enterprise-survey-2023-financial-year-provisional/Download-data/annual-enterprise-survey-2023-financial-year-provisional-size-bands.csv"
	// link = "https://www.stats.govt.nz/assets/Uploads/Business-operations-survey/Business-operations-survey-2022/Download-data/business-operations-survey-2022-business-finance.csv"

	for fileName, link := range files {
		// fileName := "survey.csv"
		filePath := tools.PrepareDir(fileName)

		tools.DownloadFile(link, filePath)
		tools.ParseFile(fileName, filePath)
	}
}
