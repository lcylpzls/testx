package core

import (
	"testing"
)

// FuzzFormatValue 验证任意输入下值格式化不 panic。
func FuzzFormatValue(f *testing.F) {
	f.Add("")
	f.Add("你好")
	f.Add("x=1&y=2")
	f.Add("\x00\x01")
	f.Fuzz(func(t *testing.T, input string) {
		_ = formatValue(input)
	})
}

// FuzzJSONEqual 验证任意 JSON 输入下语义比较不 panic。
func FuzzJSONEqual(f *testing.F) {
	f.Add("{}", "{}")
	f.Add(`{"a":1}`, `{"a":1.0}`)
	f.Add(`[1,2]`, `[1,2]`)
	f.Add("", "x")
	f.Fuzz(func(t *testing.T, a, b string) {
		JSONEqual(&fakeTB{}, a, b)
	})
}
