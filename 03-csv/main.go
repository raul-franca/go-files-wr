package main

import (
	"encoding/csv"
	"fmt"
	"os"
)

type Employee struct {
	Name    string
	Email   string
	Country string
}

var emp Employee
var employees []Employee

func ReadCSV() {
	file, err := os.Open("03-csv/demo.csv")
	if err != nil {
		fmt.Println("Error:", err)
	}
	defer file.Close()
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = 3
	reader.Comma = ';'

	csvData, err := reader.ReadAll()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	for _, item := range csvData {
		emp.Name = item[0]
		emp.Email = item[1]
		emp.Country = item[2]
		employees = append(employees, emp)
		fmt.Printf("name: %s email: %s country: %s\n", item[0], item[1], item[2])
	}
}

func WriteCSV() {
	csvFile, err := os.Create("03-csv/employees.csv")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	defer csvFile.Close()

	writer := csv.NewWriter(csvFile)
	writer.Comma = ';'

	for _, itemEmp := range employees {
		records := []string{itemEmp.Name, itemEmp.Email, itemEmp.Country}
		err := writer.Write(records)
		if err != nil {
			fmt.Println("Error:", err)
		}
	}
	writer.Flush()
}

func main() {

	ReadCSV()
	WriteCSV()
}
