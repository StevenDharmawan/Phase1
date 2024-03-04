package main

import (
	"encoding/csv"
	"log"
	"os"
	"strings"
)

func main() {
	records := readCsvFile("input.csv")
	for i := 0; i < 4; i++ {
		for j := 0; j < 3; j++ {
			if i != 0 && j == 0 {
				records[i][j] = strings.ToUpper(records[i][j])
			}
			if i != 0 && j == 2 {
				records[i][j] = "Mr." + records[i][j]
			}
		}
	}
	writeFile(records)
}

func readCsvFile(filePath string) [][]string {
	f, err := os.Open(filePath)
	if err != nil {
		log.Fatal("Unable to read input file "+filePath, err)
	}
	defer f.Close()

	csvReader := csv.NewReader(f)
	records, err := csvReader.ReadAll()
	if err != nil {
		log.Fatal("Unable to parse file as CSV for "+filePath, err)
	}

	return records
}

func writeFile(texts [][]string) {
	csvFile, err := os.Create("output.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvWriter := csv.NewWriter(csvFile)
	for _, text := range texts {
		_ = csvWriter.Write(text)
	}
	csvWriter.Flush()
}
