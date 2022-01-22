package pkg

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strings"
)

// ReadCSV opens csv file defined in path and returns parsed strings.
func ReadCSV(path string) [][]string {
	// Is the file exists?
	_, err := os.Stat(path)
	if err != nil {
		fmt.Println("Cannot stat", path)
		return nil
	}
	// Open file
	f, err := os.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		//	log.Fatal(err)
		return nil
	}
	defer f.Close()
	// Create a new CSV reader reading from the opened file.
	reader := csv.NewReader(f)

	// Assume we don't know the number of fields per line. By setting
	// FieldsPerRecord negative, each row may have a variable
	// number of fields.
	reader.FieldsPerRecord = -1

	// rawCSVData will hold our successfully parsed rows.
	var rawCSVData [][]string
	// Read in the records one by one.
	for {
		// Read in a row. Check if we are at the end of the file.
		record, err := reader.Read()
		if err == io.EOF {
			break
		}

		// Append the record to our dataset.
		rawCSVData = append(rawCSVData, strings.Split(record[0], ";"))
	}

	return rawCSVData
}
