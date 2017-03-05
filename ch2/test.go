package main

import (
	"testing"

	"GoBook/TheGoProgrammingLanguagePractice/ch2/popcount"
)

func main()  {
	// TODO 未完成
}

// -- Benchmarks基准测试 --

func BenchmarkPopCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCount(0x1234567890ABCDEF)
	}
}

func BenchmarkBitCount(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountV3(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByClearing(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountV5(0x1234567890ABCDEF)
	}
}

func BenchmarkPopCountByShifting(b *testing.B) {
	for i := 0; i < b.N; i++ {
		popcount.PopCountV4(0x1234567890ABCDEF)
	}
}