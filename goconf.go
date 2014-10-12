package goconf

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func getConfiguration(fullPath string) map[string]string {
	inputFile, err := os.Open(fullPath)
	check(err)
	defer inputFile.Close()

	var allLines map[string]string
	allLines = make(map[string]string)
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
		allLines[splitLine[0]] = splitLine[1]
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(scanner.Err())
	}

	return allLines
}

func Parse(fileName string) map[string]string {
	path, err := os.Getwd()
	check(err)
	path += ("/" + fileName)
	return getConfiguration(path)
}
