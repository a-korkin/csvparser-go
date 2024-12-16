package tools

import (
	"bufio"
	"io"
	"log"
	"net/http"
	"os"
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

func DownloadFile(uri string) {
	resp, err := http.Get(uri)
	if err != nil {
		log.Fatalf("failed to load file: %s", err)
	}
	if resp.Header.Get("Content-Type") != "text/csv" {
		log.Fatalf("file is not csv")
	}

	fo, err := os.Create("data/ex.csv")
	if err != nil {
		log.Fatalf("failed to create file: %s", err)
	}
	rw := bufio.NewReadWriter(bufio.NewReader(resp.Body), bufio.NewWriter(fo))
	writeToFile(rw)
}
