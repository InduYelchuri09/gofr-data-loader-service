package main

import (
	"fmt"
	"log"

	"your-module-path/utils" // adjust this to your module name
)

func main() {
	err := utils.CSVToJSON("sample.csv", "output.json")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("CSV to JSON done.")

	err = utils.JSONToCSV("output.json", "output.csv")
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("JSON to CSV done.")
}

