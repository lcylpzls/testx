package testx

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
