package main

import (
	"encoding/json"
	"fmt"
	"goLangBaker/app"
	"io"
	"log"
	"os"

	"github.com/sqweek/dialog"
)

func main() {

	openFilePath, err := dialog.File().Filter("JSON Files", "json").Title("Select a file to open").Load()
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Selected file:", openFilePath)

	file, err := os.Open(openFilePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	byteValue, err := io.ReadAll(file)
	if err != nil {
		log.Fatal(err)
	}

	var inputValues app.ImportValues
	err = json.Unmarshal(byteValue, &inputValues)
	if err != nil {
		log.Fatal(err)
	}

	markdownDocument := app.App(inputValues)

	saveFilePath, err := dialog.File().Filter("Markdown Files", "md").Title("Save file as").Save()
	if err != nil {
		log.Fatal(err)
	}

	// Write the markdown content to the file
	err = os.WriteFile(saveFilePath, []byte(markdownDocument), 0644)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Markdown file saved as:", saveFilePath)
}
