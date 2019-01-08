package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

func readCSV() {
	csvFile, err := os.Open("/home/shiun/Documents/Masterarbeit/Data/EEA/DataFlow_B.csv")
	if err != nil{
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	rows, err := csvReader.ReadAll()
	if err != nil{
		panic(err)
	}
	for row := range rows{
		fmt.Print(row)
	}
}
