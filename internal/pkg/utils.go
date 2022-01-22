package pkg

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"os"
)

type row []string

func readCSV(path string, out chan row) {
	csvFile, err := os.Open(path)
	if err != nil {
		log.Fatalf("Cannot open CSV: %v", err)
	}

	reader := csv.NewReader(csvFile)

	go func(reader *csv.Reader) {
		for {
			var line, error = reader.Read()
			if error == io.EOF {
				csvFile.Close()
				close(out)
				break
			} else if err != nil {
				log.Fatalf("Error reading file: %v", err)
			}
			out <- line
		}
	}(reader)
}

func CountryA2Prepare(path string) map[string]string {
	var res = make(map[string]string)

	type rawData struct {
		global_id        int
		signature_date   string
		system_object_id string
		ALFA3            string
		SHORTNAME        string
		FULLNAME         string
		ALFA2            string
		CODE             string
	}

	f, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "%s\n", err)
		//	log.Fatal(err)
	}

	//defer f.Close()

	var data []rawData
	fmt.Println(data)

	err = json.Unmarshal([]byte(f), &data)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("d -> ", data)
	return res
}
