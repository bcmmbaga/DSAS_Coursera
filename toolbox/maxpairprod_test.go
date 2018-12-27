package toolbox

import (
	"testing"
)

func TestMaxPairProd(t *testing.T) {
	sample := []struct {
		arr      []int
		expected float64
	}{
		{[]int{1, 2, 3}, 6},
		{[]int{1, 2}, 2},
		{[]int{2, 1}, 2},
		{[]int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3}, 140},
		{[]int{5, 6, 6}, 36},
		{[]int{6132, 56210, 45236, 95361, 68380, 16906, 80495, 95298}, 9087712578},
	}

	for _, v := range sample {
		if k := maxPairwiseProdNaive(v.arr); k != v.expected {
			t.Errorf("From naive func:: Expected %f got %f", v.expected, k)
		}

		if k := maxPairwiseProdFast(v.arr); k != v.expected {
			t.Errorf("From fast func:: Expected %f got %f", v.expected, k)
		}
	}
}

func BenchmarkMaxPairProdNaive(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxPairwiseProdNaive([]int{7, 5, 14, 2, 8, 8, 10, 1, 2, 3})
	}
}

func BenchmarkMaxPairProdFast(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxPairwiseProdFast([]int{6132, 56210, 45236, 95361, 68380, 16906, 80495, 95298})
	}
}

func BenchmarkMaxPairProdSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		maxPairwiseProdSort([]int{6132, 56210, 45236, 95361, 68380, 16906, 80495, 95298})
	}
}
