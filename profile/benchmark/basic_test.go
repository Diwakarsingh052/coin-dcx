package benchmark

import (
	"sort"
	"testing"
)

func BenchmarkSort1000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		sort.Ints(els)
	}
}

func BenchmarkBubbleSort1000(b *testing.B) {
	els := getElements(100000)
	for i := 0; i < b.N; i++ {
		bubbleSort(els)
	}
}
