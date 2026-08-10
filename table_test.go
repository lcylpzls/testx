package testx

import (
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
}

func TestRecoverPanic(t *testing.T) {
	if got := recoverPanic(func() { panic("x") }); got != "x" {
		t.Fatalf("recoverPanic 应返回 panic 值：%v", got)
	}
	if got := recoverPanic(func() {}); got != nil {
		t.Fatalf("recoverPanic 未 panic 应返回 nil：%v", got)
	}
}
