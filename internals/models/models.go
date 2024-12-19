package models

import (
	"log"
	"strconv"
)

type Survey struct {
	Year     uint16
	Code     string
	Name     string
	RMESize  string
	Variable string
	Value    string
	Unit     string
}

func CreateSurveys(records [][]string) {
	surveys := make([]Survey, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}
		year, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("failed to convert str to year: %s", err)
		}
		survey := Survey{
			Year:     uint16(year),
			Code:     record[1],
			Name:     record[2],
			RMESize:  record[3],
			Variable: record[4],
			Value:    record[5],
			Unit:     record[6],
		}
		surveys = append(surveys, survey)
	}
	log.Printf("count of surveys: %d\n", len(surveys))
}

type Finance struct {
	Description string
	Industry    string
	Level       string
	Size        int
	LineCode    string
	Value       int
}
