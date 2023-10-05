package main

import (
	"math"
	"strconv"
)

func calculateSTDDEV(data []string) (float64, error) {
	fData, err := parseFloat(data)
	if err != nil {
		return float64(0), err
	}
	mean := calculateMean(fData)

	sumSquaredDifferences := 0.0
	for _, num := range fData {
		sumSquaredDifferences += (num - mean) * (num - mean)
	}

	variance := sumSquaredDifferences / float64(len(fData))

	standardDeviation := math.Sqrt(variance)

	return standardDeviation, nil
}

func calculateMean(numbers []float64) float64 {
	sum := 0.0
	for _, num := range numbers {
		sum += num
	}
	return sum / float64(len(numbers))
}

func parseFloat(sNumbers []string) ([]float64, error) {
	numbers := make([]float64, len(sNumbers))
	for _, numStr := range sNumbers {
		if numStr != "" {
			num, err := strconv.ParseFloat(numStr, 64)
			if err != nil {
				return nil, err
			} else {
				numbers = append(numbers, num)
			}
		}
	}
	return numbers, nil
}
