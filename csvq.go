package main

import (
	"bufio"
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	in := bufio.NewScanner(os.Stdin)
	var records [][]string
	for in.Scan() {
		txt := in.Text()
		txt = strings.Replace(txt, `”`, `"`, -1)
		txt = strings.Replace(txt, `“`, `"`, -1)
		records = append(records, []string{txt})
	}

	w := csv.NewWriter(os.Stdout)
	for _, record := range records {
		if err := w.Write(record); err != nil {
			log.Fatalln("error writing record to csv:", err)
		}
	}

	// Write any buffered data to the underlying writer (standard output).
	w.Flush()

	if err := w.Error(); err != nil {
		log.Fatal(err)
	}
}
