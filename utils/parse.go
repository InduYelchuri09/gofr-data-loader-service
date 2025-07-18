package utils

import (
	"encoding/csv"
	"encoding/json"
	"io"
)

func ParseCSV(file io.Reader) ([]map[string]string, error) {
	reader := csv.NewReader(file)

	headers, err := reader.Read()
	if err != nil {
		return nil, err
	}

	var result []map[string]string
	for i := 0; i < 5; i++ {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			return nil, err
		}

		row := map[string]string{}
		for j, value := range record {
			row[headers[j]] = value
		}
		result = append(result, row)
	}

	return result, nil
}

func ParseJSON(file io.Reader) ([]map[string]interface{}, error) {
	decoder := json.NewDecoder(file)
	var data []map[string]interface{}

	err := decoder.Decode(&data)
	if err != nil {
		return nil, err
	}

	if len(data) > 5 {
		return data[:5], nil
	}
	return data, nil
}
