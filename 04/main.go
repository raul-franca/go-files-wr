package main

import (
	"fmt"
	"io"
	"strings"
)

type Product struct {
	Name, Category string
	Price          float64
}

var Kayak = Product{
	Name:     "Kayak",
	Category: "Watersports",
	Price:    279,
}

var Products = []Product{
	{"Kayak", "Watersports", 279},
	{"Lifejacket", "Watersports", 49.95},
	{"Soccer Ball", "Soccer", 19.50},
	{"Corner Flags", "Soccer", 34.95},
	{"Stadium", "Soccer", 79500},
	{"Thinking Cap", "Chess", 16},
	{"Unsteady Chair", "Chess", 75},
	{"Bling-Bling King", "Chess", 1200},
}

func processData(reader io.Reader, writer io.Writer) {
	b := make([]byte, 2)
	for {
		count, err := reader.Read(b)
		if count > 0 {
			writer.Write(b[0:count])
			fmt.Printf("Read %v bytes: %v ", count, string(b[0:count]))
		}
		if err == io.EOF {
			break
		}
	}
}

func main() {
	r := strings.NewReader("Kayak")
	var builder strings.Builder
	processData(r, &builder)
	fmt.Printf("String builder contents: %s", builder.String())
}
