package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"testing"
	_ "time"
)

// Take written JSON files from main and test those instead of making a new get request
// iterate through testFiles folder and return error with the name of the file that failed

var testflower ProductData

func TestProductIDs(t *testing.T) {

	// Set the directory path
	dirPath := "jsonFiles"

	// Get a list of files in the directory
	files, err := ioutil.ReadDir(dirPath)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Loop through each file in the directory
	for _, file := range files {
		// Check if the file is a regular file (not a directory or symlink)
		if !file.IsDir() && (file.Mode()&os.ModeSymlink == 0) {
			// Open the file for reading
			filePath := filepath.Join(dirPath, file.Name())
			f, err := os.Open(filePath)
			if err != nil {
				t.Fail()
				fmt.Errorf("There was an issue opening the file: " + file.Name())
				log.Fatal(err)
			}

			fmt.Println(file.Name())
			defer f.Close()

			// Do something with the file
			// For example, read its contents and print them to the console
			err = json.NewDecoder(f).Decode(&testflower)

			if err != nil {
				t.Fail()
				fmt.Errorf("Decoding Error with the json file: " + file.Name())
				log.Fatal(err)
			}
		}

	}
}
