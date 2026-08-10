package testx

import (
	"fmt"
	"math"
	"reflect"
	"testing"
)

// RunCases 以子测试运行表格用例：
// 每个用例生成名为 "name[序号]" 的子测试；fn 必须非空。
func RunCases[T any](t *testing.T, name string, cases []T, fn func(t *testing.T, c T)) {
	t.Helper()
	for i, c := range cases {
		t.Run(fmt.Sprintf("%s[%d]", name, i), func(t *testing.T) {
			t.Helper()
			fn(t, c)
		})
	}
}

// Panics 断言 fn 触发 panic。
func Panics(t TB, fn func()) {
	t.Helper()
	if recoverPanic(fn) == nil {
		t.Errorf("期望触发 panic，实际正常返回")
	}
}

// PanicsWithValue 断言 fn 触发 panic 且 panic 值与 want 深度相等。
func PanicsWithValue(t TB, want any, fn func()) {
	t.Helper()
	got := recoverPanic(fn)
	if got == nil {
		t.Errorf("期望触发 panic，实际正常返回")
		return
	}
	if !reflect.DeepEqual(got, want) {
		t.Errorf("期望 panic 值 %s，实际 %s", formatValue(want), formatValue(got))
	}
}

// NotPanics 断言 fn 不触发 panic。
func NotPanics(t TB, fn func()) {
	t.Helper()
	if got := recoverPanic(fn); got != nil {
		t.Errorf("期望不触发 panic，实际 panic 值 %s", formatValue(got))
	}
}

// Approx 断言 |got-want| <= tolerance；tolerance 必须非负。
func Approx(t TB, got, want, tolerance float64) {
	t.Helper()
	if tolerance < 0 {
		t.Fatalf("Approx 的容差必须非负，得到 %v", tolerance)
		return
	}
	if math.Abs(got-want) > tolerance {
		t.Errorf("期望 %v 与 %v 的差不超过 %v，实际差 %v",
			got, want, tolerance, math.Abs(got-want))
	}
}

// recoverPanic 执行 fn 并返回 panic 值；未 panic 时返回 nil。
func recoverPanic(fn func()) (recovered any) {
	defer func() {
		recovered = recover()
	}()
	fn()
	return nil
}
