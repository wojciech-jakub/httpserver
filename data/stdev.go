package data

import (
	"fmt"
	"math"
	"strconv"
)

func CalculateSTDEV(data []string) float64 {
	fData := parseFloat(data)
	mean := calculateMean(fData)

	sumSquaredDifferences := 0.0
	for _, num := range fData {
		sumSquaredDifferences += (num - mean) * (num - mean)
	}

	variance := sumSquaredDifferences / float64(len(fData))

	standardDeviation := math.Sqrt(variance)

	return standardDeviation
}

func calculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func parseFloat(sNumbers []string) []float64 {
	numbers := make([]float64, len(sNumbers))
	for _, numStr := range sNumbers {
		if numStr != "" {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				fmt.Println("Error parsing number:", err)
			} else {
				numbers = append(numbers, num)
			}
		}
	}
	return numbers
}
