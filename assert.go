package testx

import (
	"errors"
	"fmt"
	"reflect"
)

// Equal 断言 got 与 want 深度相等（reflect.DeepEqual）。
func Equal(t TB, got, want any) {
	t.Helper()
	if !reflect.DeepEqual(got, want) {
		t.Errorf("期望相等：\n  实际: %s\n  期望: %s", formatValue(got), formatValue(want))
	}
}

// NotEqual 断言 got 与 want 不相等。
func NotEqual(t TB, got, want any) {
	t.Helper()
	if reflect.DeepEqual(got, want) {
		t.Errorf("期望不相等，两者均为 %s", formatValue(got))
	}
}

// True 断言值为 true。
func True(t TB, v bool) {
	t.Helper()
	if !v {
		t.Errorf("期望为 true，得到 false")
	}
}

// False 断言值为 false。
func False(t TB, v bool) {
	t.Helper()
	if v {
		t.Errorf("期望为 false，得到 true")
	}
}

// Nil 断言值为 nil（含 nil 指针/切片/map/函数/通道/接口）。
func Nil(t TB, v any) {
	t.Helper()
	if !isNil(v) {
		t.Errorf("期望为 nil，得到 %s", formatValue(v))
	}
}

// NotNil 断言值非 nil。
func NotNil(t TB, v any) {
	t.Helper()
	if isNil(v) {
		t.Errorf("期望非 nil，得到 %s", formatValue(v))
	}
}

// Error 断言 err 非 nil。
func Error(t TB, err error) {
	t.Helper()
	if err == nil {
		t.Errorf("期望返回错误，得到 nil")
	}
}

// NoError 断言 err 为 nil。
func NoError(t TB, err error) {
	t.Helper()
	if err != nil {
		t.Errorf("期望无错误，得到 %s", formatValue(err))
	}
}

// ErrorIs 断言错误链中包含 target（errors.Is）。
func ErrorIs(t TB, err, target error) {
	t.Helper()
	if !errors.Is(err, target) {
		t.Errorf("期望错误链包含 %s，实际 %s", formatValue(target), formatValue(err))
	}
}

// Empty 断言值为零值/空值（空字符串、零数值、false、nil、空容器）。
func Empty(t TB, v any) {
	t.Helper()
	if !isEmpty(v) {
		t.Errorf("期望为空值，得到 %s", formatValue(v))
	}
}

// NotEmpty 断言值非零值/非空。
func NotEmpty(t TB, v any) {
	t.Helper()
	if isEmpty(v) {
		t.Errorf("期望非空值，得到 %s", formatValue(v))
	}
}

// Len 断言容器长度；支持 string/array/slice/map/chan。
// 其他类型视为断言误用，立即终止测试。
func Len(t TB, v any, want int) {
	t.Helper()
	n, ok := lengthOf(v)
	if !ok {
		t.Fatalf("Len 不支持类型 %T", v)
		return
	}
	if n != want {
		t.Errorf("期望长度 %d，实际 %d（%s）", want, n, formatValue(v))
	}
}

// isNil 判断值是否为 nil（含类型化 nil）。
func isNil(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Map, reflect.Ptr, reflect.Slice:
		return rv.IsNil()
	}
	return false
}

// isEmpty 判断值是否为零值/空值。
func isEmpty(v any) bool {
	if v == nil {
		return true
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String:
		return rv.Len() == 0
	case reflect.Bool:
		return !rv.Bool()
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return rv.Int() == 0
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return rv.Uint() == 0
	case reflect.Float32, reflect.Float64:
		return rv.Float() == 0
	case reflect.Chan, reflect.Func, reflect.Interface, reflect.Ptr:
		return rv.IsNil()
	}
	return false
}

// lengthOf 返回容器长度；不支持的类型的 ok 为 false。
func lengthOf(v any) (int, bool) {
	if v == nil {
		return 0, false
	}
	rv := reflect.ValueOf(v)
	switch rv.Kind() {
	case reflect.Array, reflect.Map, reflect.Slice, reflect.String, reflect.Chan:
		return rv.Len(), true
	}
	return 0, false
}

// formatValue 输出确定性的可读值：nil/字符串/错误特殊化，其余 %#v。
func formatValue(v any) string {
	if v == nil {
		return "nil"
	}
	if s, ok := v.(string); ok {
		return fmt.Sprintf("%q", s)
	}
	if e, ok := v.(error); ok {
		return fmt.Sprintf("error(%q)", e.Error())
	}
	return fmt.Sprintf("%#v", v)
}

// Greater 断言 got 大于 want（仅支持数值类型）。
func Greater(t TB, got, want any) {
	t.Helper()
	if cmp, ok := compareNumber(got, want); !ok {
		t.Errorf("Greater 仅支持数值类型，得到 %T 与 %T", got, want)
	} else if cmp <= 0 {
		t.Errorf("期望 %s 大于 %s", formatValue(got), formatValue(want))
	}
}

// GreaterOrEqual 断言 got 大于等于 want（仅支持数值类型）。
func GreaterOrEqual(t TB, got, want any) {
	t.Helper()
	if cmp, ok := compareNumber(got, want); !ok {
		t.Errorf("GreaterOrEqual 仅支持数值类型，得到 %T 与 %T", got, want)
	} else if cmp < 0 {
		t.Errorf("期望 %s 大于等于 %s", formatValue(got), formatValue(want))
	}
}

// Less 断言 got 小于 want（仅支持数值类型）。
func Less(t TB, got, want any) {
	t.Helper()
	if cmp, ok := compareNumber(got, want); !ok {
		t.Errorf("Less 仅支持数值类型，得到 %T 与 %T", got, want)
	} else if cmp >= 0 {
		t.Errorf("期望 %s 小于 %s", formatValue(got), formatValue(want))
	}
}

// LessOrEqual 断言 got 小于等于 want（仅支持数值类型）。
func LessOrEqual(t TB, got, want any) {
	t.Helper()
	if cmp, ok := compareNumber(got, want); !ok {
		t.Errorf("LessOrEqual 仅支持数值类型，得到 %T 与 %T", got, want)
	} else if cmp > 0 {
		t.Errorf("期望 %s 小于等于 %s", formatValue(got), formatValue(want))
	}
}

// compareNumber 比较两个数值，返回 -1/0/1；非数值返回 ok=false。
func compareNumber(a, b any) (int, bool) {
	af, aok := numericValue(a)
	bf, bok := numericValue(b)
	if !aok || !bok {
		return 0, false
	}
	switch {
	case af < bf:
		return -1, true
	case af > bf:
		return 1, true
	default:
		return 0, true
	}
}

// numericValue 将数值类型归一为 float64；非数值返回 ok=false。
func numericValue(v any) (float64, bool) {
	rv := reflect.ValueOf(v)
	if !rv.IsValid() {
		return 0, false
	}
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		return float64(rv.Int()), true
	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		return float64(rv.Uint()), true
	case reflect.Float32, reflect.Float64:
		return rv.Float(), true
	default:
		return 0, false
	}
}
