package internal

import (
	"encoding/csv"
	"os"
)

// readCSVFile reads the CSV file and returns the records slice
func readCSVFile(filename string) ([][]string, error) {
	recordFile, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	reader := csv.NewReader(recordFile)

	records, err := reader.ReadAll()
	if err != nil {
		return nil, err
	}

	err = recordFile.Close()
	if err != nil {
		return nil, err
	}

	return records, nil
}
