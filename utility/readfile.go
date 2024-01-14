package utility

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

// Function to read the ASCII Style files line by line
func ReadNumbersFile(style string) []float64 {
	readFile, err := os.Open(style)

	if err != nil {
		fmt.Println(err)
	}

	// Scan the lines that was Read into slice of strings
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	numberSlice := []float64{}
	for fileScanner.Scan() {
		// value, err := strconv.Atoi(fileScanner.Text())
		value, err := strconv.ParseFloat(fileScanner.Text(), 64)
		if err == nil {
			numberSlice = append(numberSlice, value)
		}
	}

	readFile.Close()
	return numberSlice
}
