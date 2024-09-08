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

func removeA(s string) string {
	return strings.Replace(s, "\u00a0", "", -1)
}

func ParseJson(filename, output string) {

	file, err := os.Open(filename)
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

	out, err := os.Create(output)
	if err != nil {
		log.Fatal(err)
	}
	defer out.Close()

	writer := csv.NewWriter(out)
	defer writer.Flush()

	header := []string{"Creation Date", "Location", "Room Number", "Recipient", "Description"}
	if err := writer.Write(header); err != nil {
		log.Fatal(err)
	}

	for _, r := range acts {
		var csvRow []string
		r.Description = removeA(r.Description)
		csvRow = append(csvRow, r.CreationDate, r.Location, r.RoomNumber, r.Recipient, r.Description)
		if err := writer.Write(csvRow); err != nil {
			log.Fatal(err)
		}
	}
}
