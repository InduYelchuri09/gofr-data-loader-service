package utils

import (
	"encoding/csv"
	"encoding/json"
	"os"
)

func JSONToCSV(jsonFilePath, csvFilePath string) error {
	jsonBytes, err := os.ReadFile(jsonFilePath)
	if err != nil {
		return err
	}

	var data []map[string]string
	err = json.Unmarshal(jsonBytes, &data)
	if err != nil {
		return err
	}
	if len(data) == 0 {
		return nil
	}

	csvFile, err := os.Create(csvFilePath)
	if err != nil {
		return err
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	defer writer.Flush()

	// Write headers
	headers := []string{}
	for k := range data[0] {
		headers = append(headers, k)
	}
	writer.Write(headers)

	for _, row := range data {
		record := []string{}
		for _, h := range headers {
			record = append(record, row[h])
		}
		writer.Write(record)
	}
	return nil
}
