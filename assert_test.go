package testx

import (
	"errors"
	"fmt"
	"strings"
	"testing"
)

func TestEqual(t *testing.T) {
	tb := &fakeTB{}
	Equal(tb, "你好", "你好")
	if tb.failed() {
		t.Fatalf("相等值不应失败：%v", tb.errors)
	}

	Equal(tb, "实际", "期望")
	if !tb.failed() {
		t.Fatal("不等值应失败")
	}
	if !strings.Contains(tb.errors[0], "实际") || !strings.Contains(tb.errors[0], "期望") {
		t.Fatalf("失败消息应包含实际/期望：%s", tb.errors[0])
	}
	if !tb.helperCalled {
		t.Fatal("应调用 Helper")
	}

	type point struct{ X, Y int }
	Equal(&fakeTB{}, point{1, 2}, point{1, 2})
	Equal(&fakeTB{}, []int{1, 2}, []int{1, 2})
}

func TestNotEqual(t *testing.T) {
	NotEqual(&fakeTB{}, 1, 2)
	tb := &fakeTB{}
	NotEqual(tb, "x", "x")
	if !tb.failed() {
		t.Fatal("相等值应失败")
	}
}

func TestTrueFalse(t *testing.T) {
	True(&fakeTB{}, true)
	tb := &fakeTB{}
	True(tb, false)
	if !tb.failed() {
		t.Fatal("false 应失败")
	}
	False(&fakeTB{}, false)
	tb2 := &fakeTB{}
	False(tb2, true)
	if !tb2.failed() {
		t.Fatal("true 应失败")
	}
}

func TestNilNotNil(t *testing.T) {
	var p *int
	var s []int
	var m map[string]int
	Nil(&fakeTB{}, nil)
	Nil(&fakeTB{}, p)
	Nil(&fakeTB{}, s)
	Nil(&fakeTB{}, m)
	tb := &fakeTB{}
	Nil(tb, 42)
	if !tb.failed() {
		t.Fatal("非 nil 应失败")
	}
	NotNil(&fakeTB{}, 42)
	tb2 := &fakeTB{}
	NotNil(tb2, p)
	if !tb2.failed() {
		t.Fatal("nil 应失败")
	}
}

func TestErrorNoError(t *testing.T) {
	sentinel := errors.New("失败")
	Error(&fakeTB{}, sentinel)
	tb := &fakeTB{}
	Error(tb, nil)
	if !tb.failed() {
		t.Fatal("nil 错误应失败")
	}
	NoError(&fakeTB{}, nil)
	tb2 := &fakeTB{}
	NoError(tb2, sentinel)
	if !tb2.failed() {
		t.Fatal("非 nil 错误应失败")
	}
	if !strings.Contains(tb2.errors[0], "失败") {
		t.Fatalf("错误消息应包含错误文本：%s", tb2.errors[0])
	}
}

func TestErrorIs(t *testing.T) {
	sentinel := errors.New("哨兵")
	wrapped := fmt.Errorf("包装：%w", sentinel)
	ErrorIs(&fakeTB{}, wrapped, sentinel)
	tb := &fakeTB{}
	ErrorIs(tb, errors.New("其他"), sentinel)
	if !tb.failed() {
		t.Fatal("不匹配应失败")
	}
}

func TestEmptyNotEmpty(t *testing.T) {
	Empty(&fakeTB{}, "")
	Empty(&fakeTB{}, 0)
	Empty(&fakeTB{}, uint(0))
	Empty(&fakeTB{}, float64(0))
	Empty(&fakeTB{}, false)
	Empty(&fakeTB{}, nil)
	Empty(&fakeTB{}, []int{})
	Empty(&fakeTB{}, [0]int{})
	Empty(&fakeTB{}, (*int)(nil))
	Empty(&fakeTB{}, (chan int)(nil))
	Empty(&fakeTB{}, (func())(nil))
	tb := &fakeTB{}
	Empty(tb, "x")
	if !tb.failed() {
		t.Fatal("非空应失败")
	}
	NotEmpty(&fakeTB{}, "x")
	NotEmpty(&fakeTB{}, uint(1))
	NotEmpty(&fakeTB{}, float64(1.5))
	NotEmpty(&fakeTB{}, &struct{}{})
	NotEmpty(&fakeTB{}, struct{ A int }{1})
	tb2 := &fakeTB{}
	NotEmpty(tb2, "")
	if !tb2.failed() {
		t.Fatal("空值应失败")
	}
}

func TestLen(t *testing.T) {
	Len(&fakeTB{}, "你好", 2)
	Len(&fakeTB{}, []int{1, 2, 3}, 3)
	Len(&fakeTB{}, map[string]int{"a": 1}, 1)
	Len(&fakeTB{}, [2]int{1, 2}, 2)
	ch := make(chan int, 3)
	Len(&fakeTB{}, ch, 3)
	tb := &fakeTB{}
	Len(tb, "你好", 3)
	if !tb.failed() {
		t.Fatal("长度不匹配应失败")
	}
	tb2 := &fakeTB{}
	Len(tb2, 42, 1)
	if len(tb2.fatals) != 1 {
		t.Fatalf("误用应 Fatalf：%v", tb2.fatals)
	}
	tb3 := &fakeTB{}
	Len(tb3, nil, 0)
	if len(tb3.fatals) != 1 {
		t.Fatalf("nil 应 Fatalf：%v", tb3.fatals)
	}
}

func TestFormatValue(t *testing.T) {
	if got := formatValue(nil); got != "nil" {
		t.Fatalf("nil 格式化不匹配：%q", got)
	}
	if got := formatValue("x"); got != `"x"` {
		t.Fatalf("字符串格式化不匹配：%q", got)
	}
	if got := formatValue(errors.New("错")); got != `error("错")` {
		t.Fatalf("错误格式化不匹配：%q", got)
	}
	if got := formatValue(struct{ A int }{1}); !strings.Contains(got, "A") {
		t.Fatalf("结构体格式化不匹配：%q", got)
	}
}

// fakeTB 是 TB 的测试替身，记录失败而不终止。
type fakeTB struct {
	helperCalled bool
	errors       []string
	fatals       []string
}

func (f *fakeTB) Helper() { f.helperCalled = true }

func (f *fakeTB) Errorf(format string, args ...any) {
	f.errors = append(f.errors, fmt.Sprintf(format, args...))
}

func (f *fakeTB) Fatalf(format string, args ...any) {
	f.fatals = append(f.fatals, fmt.Sprintf(format, args...))
}

func (f *fakeTB) Cleanup(func()) {}

func (f *fakeTB) failed() bool {
	return len(f.errors) > 0 || len(f.fatals) > 0
}
