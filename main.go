package main

import (
	"fmt"
	"main/myMath"
	"main/utility"
	"math"
	"os"
)

func main() {
	if len(os.Args) == 2 {
		fileName := os.Args[1]
		numbers := utility.ReadNumbersFile(fileName)

		average := myMath.Average(numbers)
		fmt.Println("Average", math.Round(average))
		median := myMath.Median(numbers)
		fmt.Println("Median", math.Round(median))
		variance := myMath.Variance(numbers, average)
		fmt.Println("Variance", math.Round(variance))
		std := myMath.StandardDeviation(variance)
		fmt.Println("Standard Deviation", math.Round(std))
	} else {
		fmt.Println("Please add the file name")
		fmt.Println("Usage: go run . your_file.txt")
	}
}
