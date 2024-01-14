package myMath

import "sort"

func Median(numbers []float64) float64 {
	sortedNumbers := numbers
	sort.Slice(sortedNumbers, func(i, j int) bool {
		return sortedNumbers[i] < sortedNumbers[j]
	})
	if len(numbers)%2 == 1 {
		return float64(sortedNumbers[len(numbers)/2])
	} else {
		return (float64(sortedNumbers[len(numbers)/2]) + float64(sortedNumbers[(len(numbers)/2)-1])) / 2.0
	}

}
