package utils

import (
	"encoding/csv"
	"encoding/json"
	"io"
	"os"
)

func CSVToJSON(csvFilePath, jsonFilePath string) error {
	csvFile, err := os.Open(csvFilePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	reader := csv.NewReader(csvFile)
	headers, err := reader.Read()
	if err != nil {
		return err
	}

	var data []map[string]string
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return err
		}
		row := make(map[string]string)
		for i, val := range record {
			row[headers[i]] = val
		}
		data = append(data, row)
	}

	jsonBytes, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}
	return os.WriteFile(jsonFilePath, jsonBytes, 0644)
}
