package testx

import (
	"strings"
	"testing"
)

func TestContains(t *testing.T) {
	Contains(&fakeTB{}, "你好世界", "世界")
	Contains(&fakeTB{}, []int{1, 2, 3}, 2)
	Contains(&fakeTB{}, [2]string{"a", "b"}, "b")
	Contains(&fakeTB{}, map[string]int{"k": 1}, "k")

	tb := &fakeTB{}
	Contains(tb, "你好世界", "火星")
	if !tb.failed() {
		t.Fatal("子串不匹配应失败")
	}
	tb2 := &fakeTB{}
	Contains(tb2, []int{1, 2}, 9)
	if !tb2.failed() {
		t.Fatal("切片元素不匹配应失败")
	}
	tb3 := &fakeTB{}
	Contains(tb3, map[string]int{"k": 1}, "x")
	if !tb3.failed() {
		t.Fatal("map 键不匹配应失败")
	}
	tb4 := &fakeTB{}
	Contains(tb4, "你好", 42)
	if len(tb4.fatals) != 1 {
		t.Fatalf("字符串容器元素类型错误应 Fatalf：%v", tb4.fatals)
	}
	tb5 := &fakeTB{}
	Contains(tb5, 42, 42)
	if len(tb5.fatals) != 1 {
		t.Fatalf("不支持容器应 Fatalf：%v", tb5.fatals)
	}
}

func TestNotContains(t *testing.T) {
	NotContains(&fakeTB{}, "你好世界", "火星")
	NotContains(&fakeTB{}, []int{1, 2}, 9)
	NotContains(&fakeTB{}, map[string]int{"k": 1}, "x")

	tb := &fakeTB{}
	NotContains(tb, "你好世界", "世界")
	if !tb.failed() {
		t.Fatal("包含时应失败")
	}
	tb2 := &fakeTB{}
	NotContains(tb2, []int{1, 2}, 2)
	if !tb2.failed() {
		t.Fatal("包含时应失败")
	}
	tb3 := &fakeTB{}
	NotContains(tb3, map[string]int{"k": 1}, "k")
	if !tb3.failed() {
		t.Fatal("包含键时应失败")
	}
	tb4 := &fakeTB{}
	NotContains(tb4, "你好", 42)
	if len(tb4.fatals) != 1 {
		t.Fatalf("字符串容器元素类型错误应 Fatalf：%v", tb4.fatals)
	}
	tb5 := &fakeTB{}
	NotContains(tb5, 42, 42)
	if len(tb5.fatals) != 1 {
		t.Fatalf("不支持容器应 Fatalf：%v", tb5.fatals)
	}
}

func TestSubset(t *testing.T) {
	Subset(&fakeTB{}, []int{1, 2, 3}, []int{3, 1})
	Subset(&fakeTB{}, []int{1, 1, 2}, []int{1, 1})
	Subset(&fakeTB{}, []int{1, 2}, []int{})
	Subset(&fakeTB{}, []int{}, []int{})

	tb := &fakeTB{}
	Subset(tb, []int{1, 2}, []int{2, 2})
	if !tb.failed() {
		t.Fatal("多重性不足应失败")
	}
	tb2 := &fakeTB{}
	Subset(tb2, 42, []int{1})
	if len(tb2.fatals) != 1 {
		t.Fatalf("不支持类型应 Fatalf：%v", tb2.fatals)
	}
	tb3 := &fakeTB{}
	Subset(tb3, []int{1}, 42)
	if len(tb3.fatals) != 1 {
		t.Fatalf("不支持类型应 Fatalf：%v", tb3.fatals)
	}
	tb4 := &fakeTB{}
	Subset(tb4, nil, []int{1})
	if len(tb4.fatals) != 1 {
		t.Fatalf("nil 列表应 Fatalf：%v", tb4.fatals)
	}
}

func TestElementsMatch(t *testing.T) {
	ElementsMatch(&fakeTB{}, []int{1, 2, 3}, []int{3, 1, 2})
	ElementsMatch(&fakeTB{}, []string{"a", "a", "b"}, []string{"b", "a", "a"})

	tb := &fakeTB{}
	ElementsMatch(tb, []int{1, 2}, []int{1, 2, 3})
	if !tb.failed() {
		t.Fatal("长度不同应失败")
	}
	if !strings.Contains(tb.errors[0], "A:") || !strings.Contains(tb.errors[0], "B:") {
		t.Fatalf("失败消息应包含 A/B：%s", tb.errors[0])
	}
	tb2 := &fakeTB{}
	ElementsMatch(tb2, []int{1, 1}, []int{1, 2})
	if !tb2.failed() {
		t.Fatal("元素不同应失败")
	}
	tb3 := &fakeTB{}
	ElementsMatch(tb3, 42, []int{1})
	if len(tb3.fatals) != 1 {
		t.Fatalf("不支持类型应 Fatalf：%v", tb3.fatals)
	}
	tb4 := &fakeTB{}
	ElementsMatch(tb4, []int{1}, "x")
	if len(tb4.fatals) != 1 {
		t.Fatalf("不支持类型应 Fatalf：%v", tb4.fatals)
	}
	tb5 := &fakeTB{}
	ElementsMatch(tb5, nil, nil)
	if len(tb5.fatals) != 1 {
		t.Fatalf("nil 列表应 Fatalf：%v", tb5.fatals)
	}
}
