package main

import (
	"log"
	"os"
	"strings"
)

// iscvs.go is a set of code which is meant to read the directory
// In reading the directory, we are looking for a csv filetype.

func ReadDir() string {
	var filename string
	files, err := os.ReadDir(".")
	if err != nil {
		log.Fatal(err)
	}
	for _, file := range files {
		switch {
		case strings.Contains(file.Name(), ".json"):
			filename = file.Name()
		case strings.Contains(file.Name(), ".csv"):
			filename = file.Name()
		}
	}
	return filename
}
