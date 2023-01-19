package test

import (
	"math/rand"
	"testing"
	"time"
)

func BenchmarkInline(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10000000000)
	y := rand.Intn(10000000000)
	for i := 0; i < b.N; i++ {
		addInline(x, y)
	}

}
func addInline(a, b int) int {
	return a + b
}

func BenchmarkNoInline(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	x := rand.Intn(10000000000)
	y := rand.Intn(10000000000)
	for i := 0; i < b.N; i++ {
		addNoInline(x, y)
	}

}

//go:noinline
func addNoInline(a, b int) int {
	return a + b
}
