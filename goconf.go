package goconf

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func getConfiguration(fullPath string) map[string]interface{} {
	inputFile, err := os.Open(fullPath)
	if err != nil {
		panic(err)
	}
	defer inputFile.Close()

	var allLines map[string]interface{}
	allLines = make(map[string]interface{})
	scanner := bufio.NewScanner(inputFile)

	for scanner.Scan() {
		currentLine := scanner.Text()
		if len(currentLine) == 0 {
			continue
		}
		splitLine := strings.Split(currentLine, "=")
		for i := 0; i < len(splitLine); i++ {
			splitLine[i] = strings.Trim(splitLine[i], " ")
		}
		if splitLineConverted, err := strconv.Atoi(splitLine[1]); err != nil {
			allLines[splitLine[0]] = splitLine[1]
		} else {
			allLines[splitLine[0]] = splitLineConverted
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}

	return allLines
}

func Parse(fileName string) map[string]interface{} {
	path, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	path += ("/" + fileName)
	return getConfiguration(path)
}
