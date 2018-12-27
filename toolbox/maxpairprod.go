package toolbox

import (
	"math"
	"sort"
)

func maxPairwiseProdNaive(numbers []int) float64 {
	var product float64

	for i := 0; i < len(numbers); i++ {
		for j := i + 1; j < len(numbers); j++ {
			product = math.Max(product, float64(numbers[i]*numbers[j]))
		}
	}

	return product
}

func maxPairwiseProdFast(numbers []int) float64 {
	var index1, index2 int

	for i := 1; i < len(numbers); i++ {
		if numbers[i] > numbers[index1] {
			index1 = i
		}
	}

	if index1 == 0 {
		index2 = 1
	}

	for i := 1; i < len(numbers); i++ {
		if i != index1 && numbers[i] > numbers[index2] {
			index2 = i
		}
	}

	return float64(numbers[index1] * numbers[index2])
}

func maxPairwiseProdSort(numbers []int) float64 {
	sort.Ints(numbers)
	n := len(numbers)

	return float64(numbers[n-2] * numbers[n-1])
}
