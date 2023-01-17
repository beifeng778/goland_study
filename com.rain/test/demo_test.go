package test

import (
	"bytes"
	"strings"
	"testing"
)

func Plus(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func StrBuilder(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func ByteBuffer(n int, str string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(str)
	}
	return buf.String()
}

func BenchmarkPlus(b *testing.B) {
	str := "hello"
	s := ""
	for i := 0; i < b.N; i++ {
		s += str
	}
}

func BenchmarkStrBuilder(b *testing.B) {
	var builder strings.Builder
	str := "hello"
	for i := 0; i < b.N; i++ {
		builder.WriteString(str)
	}
}

func BenchmarkByteBuffer(b *testing.B) {
	str := "hello"
	buf := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		buf.WriteString(str)
	}
}

func BenchmarkEmptyStructMap(b *testing.B) {
	m := make(map[int]struct{})
	for i := 0; i < b.N; i++ {
		m[i] = struct{}{}
	}
}

func BenchmarkBoolMap(b *testing.B) {
	m := make(map[int]bool)
	for i := 0; i < b.N; i++ {
		m[i] = false
	}
}
