package ringo

import (
	"math/rand"
	"testing"
)

func BenchmarkF64Append(b *testing.B) {
	buf := CircleF64{}
	buf.Init(100)
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Append(rand.Float64())
	}
}

func BenchmarkF64ContiguousLargeBufferFilled(b *testing.B) {
	size := 10000
	buf := CircleF64{}
	buf.Init(size)
	for idx := 0; idx < size+(size/2); idx++ {
		buf.Append(rand.Float64())
	}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		buf.Contiguous()
	}
}
