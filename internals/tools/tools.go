package tools

import (
	"bufio"
	"encoding/csv"
	"github.com/a-korkin/csvparser/internals/models"
	"io"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
)

func writeToFile(rw *bufio.ReadWriter) {
	buf := make([]byte, 1024)
	for {
		n, err := rw.Reader.Read(buf)
		if err != nil && err != io.EOF {
			log.Fatalf("failed to read body: %s", err)
		}
		if n == 0 {
			break
		}
		n, err = rw.Writer.Write(buf[:n])
		if err != nil {
			log.Fatalf("failed to write buf in file: %s", err)
		}
		err = rw.Writer.Flush()
		if err != nil {
			log.Fatalf("failed flush to file: %s", err)
		}
	}
}

func PrepareDir(fileName string) string {
	dirName := os.Getenv("DIR_NAME")
	if err := os.Mkdir(dirName, os.ModePerm); err != nil {
		log.Fatalf("failed to create directory: %s", err)
	}
	return filepath.Join(dirName, fileName)
}

func DownloadFile(uri string, fileName string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("failed to load file: %s", err)
	}
	if resp.Header.Get("Content-Type") != "text/csv" {
		log.Fatalf("file is not csv")
	}

	fo, err := os.Create(fileName)
	defer func() {
		err := fo.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}()
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	rw := bufio.NewReadWriter(bufio.NewReader(resp.Body), bufio.NewWriter(fo))
	writeToFile(rw)
}

func ParseFile(filePath string) {
	fi, err := os.Open(filePath)
	if err != nil {
		log.Fatalf("failed to open file: %s", err)
	}
	defer func() {
		err := fi.Close()
		if err != nil {
			log.Fatalf("failed to close file: %s", err)
		}
	}()
	reader := csv.NewReader(fi)
	records, err := reader.ReadAll()
	reader.LazyQuotes = true
	if err != nil {
		log.Fatalf("failed to read from file: %s", err)
	}
	surveys := make([]models.Survey, 0)
	for i, record := range records {
		if i == 0 {
			continue
		}
		year, err := strconv.Atoi(record[0])
		if err != nil {
			log.Fatalf("failed to convert str to year: %s", err)
		}
		survey := models.Survey{
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
