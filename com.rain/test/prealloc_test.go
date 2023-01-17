package test

import "testing"

func BenchmarkPreAlloc(b *testing.B) {
	data := make([]int, 0, b.N)
	for i := 0; i < b.N; i++ {
		data = append(data, i)
	}
}
