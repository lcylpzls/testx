package core

import (
	"testing"
)

// BenchmarkEqual 基准：字符串深度相等断言。
func BenchmarkEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Equal(&silentTB{}, "你好，世界", "你好，世界")
	}
}

// BenchmarkJSONEqual 基准：JSON 语义相等。
func BenchmarkJSONEqual(b *testing.B) {
	for i := 0; i < b.N; i++ {
		JSONEqual(&silentTB{}, `{"a":1,"b":[1,2]}`, `{"b":[1,2],"a":1.0}`)
	}
}

// BenchmarkContains 基准：切片元素包含断言。
func BenchmarkContains(b *testing.B) {
	data := []int{1, 2, 3, 4, 5, 6, 7, 8}
	for i := 0; i < b.N; i++ {
		Contains(&silentTB{}, data, 8)
	}
}

// silentTB 是基准专用替身：断言应全部通过，不产生输出。
type silentTB struct{}

func (silentTB) Helper()                           {}
func (silentTB) Errorf(format string, args ...any) {}
func (silentTB) Fatalf(format string, args ...any) {}
func (silentTB) Cleanup(func())                    {}
