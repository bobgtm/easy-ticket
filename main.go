package main

import (
	"bytes"
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strings"
)

func removeBOM(data []byte) []byte {
	bom := []byte{0xEF, 0xBB, 0xBF}
	if bytes.HasPrefix(data, bom) {
		return data[len(bom):]
	}
	return data
}

func removeAh(s string) string {
	return strings.Replace(s, "\u00a0", "", -1)
}

func main() {
	file, err := os.Open("myactions.json")
	if err != nil {
		log.Fatal(err)
	}
	data, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("error reading file:", err)
	}

	cleaned := removeBOM(data)
	var acts []Action
	err = json.Unmarshal(cleaned, &acts)
	if err != nil {
		log.Fatal(err)
	}

	output, err := os.Create("act.csv")
	if err != nil {
		log.Fatal(err)
	}
	defer output.Close()

	writer := csv.NewWriter(output)
	defer writer.Flush()

	header := []string{"Creation Date", "Location", "Room Number", "Recipient", "Description"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}
	// var c = []byte("Â")

	for _, r := range acts {
		var csvRow []string
		r.Description = removeAh(r.Description)
		csvRow = append(csvRow, r.CreationDate, r.Location, r.RoomNumber, r.Recipient, r.Description)
		if err := writer.Write(csvRow); err != nil {
			log.Fatal(err)
		}
	}
}
