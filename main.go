package main

import (
	"encoding/csv"
	"fmt"
	"log"

	//"io"
	//"io/ioutil"
	//"log"
	"os"
	//"strings"
)

func readCSV2() [][]string {

	csvFile, err := os.Open("/home/shiun/Documents/Masterarbeit/Data/ReWriteTest.csv")
	//csvFile, err := os.Open("/home/shiun/Documents/Masterarbeit/Data/FinalForm.csv")
	if err != nil {
		panic(err)
	}
	defer csvFile.Close()

	csvReader := csv.NewReader(csvFile)

	//row, err := csvReader.Read()
	//fmt.Print(row)

	rows, err := csvReader.ReadAll()
	if err != nil {
		panic(err)
	}

	len := len(rows)
	//fmt.Println(rows[0][])
	for i, row := range rows {
		fmt.Print(i, row[:], "\n")
	}

	fmt.Print(len, "\n")

	/*
	// i is the number of the column
	for i, row := range rows{
		fmt.Print(i, row, "\n")
	}*/

	return rows

	/*
	csvFile, err := ioutil.ReadFile("/home/shiun/Documents/Masterarbeit/Data/EEA/DataFlow_B.csv")
	if err != nil {
		panic(err)
	}

	row := csv.NewReader(strings.NewReader(string((csvFile[:]))))

	for{
		record, err := row.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		fmt.Print(record, "\n")
	}

	return row*/
}

func writeCSV(data [][]string) {
	csvFile, err := os.Create("/home/shiun/Documents/Masterarbeit/Data/ReWriteTest.csv")
	if err != nil {
		panic(err)
	}

	csvFile.WriteString("\xEF\xBB\xBF") //UTF-8

	writeFile := csv.NewWriter(csvFile)
	/*firstRow := [][]string{
		{"Country", "City", "Time", "PM2.5", "PM10", "O3"},
	}

	writeFile.WriteAll(firstRow)
*/
	writeFile.WriteAll(data)
	writeFile.Flush()
}

func writeLastLine() {
	oriData := readCSV2()

	csvFile, err := os.Create("/home/shiun/Documents/Masterarbeit/Data/ReWriteTest.csv")
	checkError(err)
	defer csvFile.Close()

	csvFile.WriteString("\xEF\xBB\xBF") //UTF-8

	data := []string{
		"A", "B", "C", "D", "E", "F", "G", "H", "I", "J", "K", "L", "M", "N", "O", "P", "Q", "R", "S", "T",
	}

	oriData = append(oriData, data)

	writeFile := csv.NewWriter(csvFile)
	//err = writeFile.Write(data)
	err = writeFile.WriteAll(oriData)
	checkError(err)

	/*
	writeFile.WriteAll(data)
	err = writeFile.WriteAll(data)
	if err != nil {
		log.Panic(err)
	}
	*/

	writeFile.Flush()
	err = writeFile.Error()
	checkError(err)
}

func checkError(err error) {
	if err != nil{
		log.Fatal(err)
	}
}

func main() {
	//data := readCSV2()
	//writeCSV(data)ReWriteTest
	//readCSV2()
	writeLastLine()
	readCSV2()
}
