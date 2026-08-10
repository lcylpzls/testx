package core

import (
	"math"
	"strings"
	"testing"
)

func TestRunCases(t *testing.T) {
	cases := []struct {
		name string
		n    int
	}{
		{"一", 1},
		{"二", 2},
	}
	ran := 0
	RunCases(t, "求和", cases, func(t *testing.T, c struct {
		name string
		n    int
	}) {
		ran++
		True(t, c.n > 0)
	})
	if ran != 2 {
		t.Fatalf("期望运行 2 个用例，实际 %d", ran)
	}
	RunCases(t, "空表格", []int{}, func(t *testing.T, c int) {})
}

func TestPanics(t *testing.T) {
	Panics(&fakeTB{}, func() { panic("崩溃") })
	tb := &fakeTB{}
	Panics(tb, func() {})
	if !tb.failed() {
		t.Fatal("未 panic 应失败")
	}
}

func TestPanicsWithValue(t *testing.T) {
	PanicsWithValue(&fakeTB{}, "崩溃", func() { panic("崩溃") })
	tb := &fakeTB{}
	PanicsWithValue(tb, "期望值", func() { panic("实际值") })
	if !tb.failed() {
		t.Fatal("panic 值不同应失败")
	}
	if !strings.Contains(tb.errors[0], "期望 panic 值") {
		t.Fatalf("失败消息缺失：%s", tb.errors[0])
	}
	tb2 := &fakeTB{}
	PanicsWithValue(tb2, "x", func() {})
	if !tb2.failed() {
		t.Fatal("未 panic 应失败")
	}
	tb3 := &fakeTB{}
	Panics(tb3, func() { panic(nil) })
	if tb3.failed() {
		t.Fatalf("panic(nil) 应视为已触发：%v", tb3.errors)
	}
}

func TestNotPanics(t *testing.T) {
	NotPanics(&fakeTB{}, func() {})
	tb := &fakeTB{}
	NotPanics(tb, func() { panic("崩溃") })
	if !tb.failed() {
		t.Fatal("panic 应失败")
	}
}

func TestApprox(t *testing.T) {
	Approx(&fakeTB{}, 1.0, 1.05, 0.1)
	Approx(&fakeTB{}, 1.0, 1.0, 0)
	tb := &fakeTB{}
	Approx(tb, 1.0, 2.0, 0.1)
	if !tb.failed() {
		t.Fatal("超出容差应失败")
	}
	if !strings.Contains(tb.errors[0], "实际差 1") {
		t.Fatalf("失败消息应包含实际差：%s", tb.errors[0])
	}
	tb2 := &fakeTB{}
	Approx(tb2, 1.0, 1.0, -1)
	if len(tb2.fatals) != 1 {
		t.Fatalf("负容差应 Fatalf：%v", tb2.fatals)
	}
	tb3 := &fakeTB{}
	Approx(tb3, NaN(), 1.0, 0.1)
	if !tb3.failed() {
		t.Fatal("NaN 实际值应失败")
	}
	tb4 := &fakeTB{}
	Approx(tb4, 1.0, NaN(), 0.1)
	if !tb4.failed() {
		t.Fatal("NaN 期望值应失败")
	}
}

func TestRecoverPanic(t *testing.T) {
	got, panicked := recoverPanic(func() { panic("x") })
	if got != "x" || !panicked {
		t.Fatalf("recoverPanic 应返回 panic 值：%v", got)
	}
	got, panicked = recoverPanic(func() {})
	if got != nil || panicked {
		t.Fatalf("recoverPanic 未 panic 应返回 nil：%v", got)
	}
	if _, panicked := recoverPanic(func() { panic(nil) }); !panicked {
		t.Fatal("panic(nil) 应标记为已触发")
	}
}

// NaN 返回 float64 的 NaN 值。
func NaN() float64 {
	return math.NaN()
}
