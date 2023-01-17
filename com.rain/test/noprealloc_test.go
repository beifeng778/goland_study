package test

import "testing"

func BenchmarkNoPreAlloc(b *testing.B) {
	data := make([]int, 0)
	for i := 0; i < b.N; i++ {
		data = append(data, i)
	}
}
