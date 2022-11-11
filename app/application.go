package app

import (
	"errors"
	"io/ioutil"
	"log"
	"os"
	"pepperstone/csvprocess"
	"strings"
)

// Run ...
func Run() {

	// Grab all csv files
	files, err := ReadFromDirectory()
	ch := make(chan bool)

	if err != nil {
		log.Println(err)
	}

	for _, file := range files {
		// Get csv file type
		fileType, err := GetFileType(file)

		if err != nil {
			log.Println(err)
		}

		// Determine the type and process the csv files
		switch fileType {
		case "ctt": // Process Cash Transfer Type CSV
			log.Println("Cash Transfer Type CSV has now been generated")
			go csvprocess.ProcesCTT(file, ch)
			<-ch
		default:
		}
	}

}

// ReadFromDirectory ...
func ReadFromDirectory() ([]string, error) {
	var csvFiles []string
	dir, err := os.Getwd()
	if err != nil {
		return nil, errors.New("Invalid directory")
	}

	files, err := ioutil.ReadDir(dir + "/csv/input")
	if err != nil {
		return nil, errors.New("Invalid directory")
	}

	for _, file := range files {
		csvFiles = append(csvFiles, file.Name())
	}

	return csvFiles, nil
}

// GetFileType ...
func GetFileType(file string) (string, error) {
	r := strings.Split(file, "_")

	if len(r) != 2 {
		return "", errors.New("Invalid filename format")
	}

	return strings.ReplaceAll(r[1], ".csv", ""), nil
}
