package goafm

import (
	"strconv"
	"testing"
)

// ~ 9.5 ns/op
func BenchmarkStringMapAssign(b *testing.B) {
	type obj struct {
		field1 int
		field2 int
	}
	keys := make([]string, 0, 500)
	for i := 0; i < 500; i++ {
		keys = append(keys, strconv.Itoa(i))
	}
	m := make(map[string]obj)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m["key"] = obj{field1: 1, field2: 2}
	}
	b.ReportAllocs()
}

// ~ 5214 ns/op
func BenchmarkStringMapRead(b *testing.B) {
	type obj struct {
		field1 int
		field2 int
	}
	m := make(map[string]obj)
	keys := make([]string, 0, 500)
	for i := 0; i < 500; i++ {
		s := strconv.Itoa(i)
		keys = append(keys, s)
		m[s] = obj{field1: 1, field2: 2}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			_ = m[k]
		}
	}
	b.ReportAllocs()
}

// ~ 7 ns/op
func BenchmarkIntMapAssign(b *testing.B) {
	type obj struct {
		field1 int
		field2 int
	}
	keys := make([]int, 0, 500)
	for i := 0; i < 500; i++ {
		keys = append(keys, i)
	}
	m := make(map[int]obj)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		m[1] = obj{field1: 1, field2: 2}
	}
	b.ReportAllocs()
}

// ~ 4495 ns/op
func BenchmarkIntMapRead(b *testing.B) {
	type obj struct {
		field1 int
		field2 int
	}
	m := make(map[int]obj)
	keys := make([]int, 0, 500)
	for i := 0; i < 500; i++ {
		keys = append(keys, i)
		m[i] = obj{field1: 1, field2: 2}
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		for _, k := range keys {
			_ = m[k]
		}
	}
	b.ReportAllocs()
}
