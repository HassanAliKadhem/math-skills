package myMath

func Variance(numbers []float64, average float64) float64 {
	differences := []float64{}

	for _, number := range numbers {
		differences = append(differences, number-average)
	}

	var sum float64 = 0

	for _, difference := range differences {
		sum += difference * difference
	}
	return sum / float64(len(numbers))
}
