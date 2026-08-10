package testx

import (
	"strings"
	"testing"
)

func TestJSONEqualPass(t *testing.T) {
	JSONEqual(&fakeTB{}, `{"a":1,"b":[1,2]}`, `{"b":[1,2],"a":1}`)
	JSONEqual(&fakeTB{}, `{"n":1}`, `{"n":1.0}`)
	JSONEqual(&fakeTB{}, `{"big":9007199254740993}`, `{"big":9007199254740993}`)
	JSONEqual(&fakeTB{}, `{"n":1e5}`, `{"n":100000}`)
	JSONEqual(&fakeTB{}, []byte(`true`), `true`)
	JSONEqual(&fakeTB{}, `"str"`, `"str"`)
	JSONEqual(&fakeTB{}, `null`, `null`)
}

func TestJSONEqualFail(t *testing.T) {
	tb := &fakeTB{}
	JSONEqual(tb, `{"a":1}`, `{"a":2}`)
	if !tb.failed() {
		t.Fatal("值不同应失败")
	}
	if !strings.Contains(tb.errors[0], "实际") || !strings.Contains(tb.errors[0], "期望") {
		t.Fatalf("失败消息应包含实际/期望：%s", tb.errors[0])
	}
	tb2 := &fakeTB{}
	JSONEqual(tb2, `{"a":1}`, `{"b":1}`)
	if !tb2.failed() {
		t.Fatal("键不同应失败")
	}
	tb3 := &fakeTB{}
	JSONEqual(tb3, `[1,2]`, `[1,2,3]`)
	if !tb3.failed() {
		t.Fatal("数组长度不同应失败")
	}
	tb4 := &fakeTB{}
	JSONEqual(tb4, `{"a":"1"}`, `{"a":1}`)
	if !tb4.failed() {
		t.Fatal("字符串与数字应视为不同")
	}
	tb4b := &fakeTB{}
	JSONEqual(tb4b, `{"a":1}`, `{"a":"1"}`)
	if !tb4b.failed() {
		t.Fatal("数字与字符串应视为不同")
	}
	tb5 := &fakeTB{}
	JSONEqual(tb5, `{"a":1,"b":2}`, `{"a":1}`)
	if !tb5.failed() {
		t.Fatal("map 长度不同应失败")
	}
	tb6 := &fakeTB{}
	JSONEqual(tb6, `[1]`, `{"0":1}`)
	if !tb6.failed() {
		t.Fatal("数组与其他类型应视为不同")
	}
	tb6b := &fakeTB{}
	JSONEqual(tb6b, `{"a":1}`, `[1]`)
	if !tb6b.failed() {
		t.Fatal("对象与其他类型应视为不同")
	}
	tb6c := &fakeTB{}
	JSONEqual(tb6c, `[1,2]`, `[1,3]`)
	if !tb6c.failed() {
		t.Fatal("数组元素不同应失败")
	}
	tb7 := &fakeTB{}
	JSONEqual(tb7, `1e5`, `2e5`)
	if !tb7.failed() {
		t.Fatal("科学计数法数值不同应失败")
	}
}

func TestJSONEqualInvalid(t *testing.T) {
	tb := &fakeTB{}
	JSONEqual(tb, `{bad}`, `{}`)
	if len(tb.fatals) != 1 {
		t.Fatalf("非法 JSON 应 Fatalf：%v", tb.fatals)
	}
	tb2 := &fakeTB{}
	JSONEqual(tb2, `{}`, `{bad}`)
	if len(tb2.fatals) != 1 {
		t.Fatalf("非法 JSON 应 Fatalf：%v", tb2.fatals)
	}
	tb3 := &fakeTB{}
	JSONEqual(tb3, `{} {}`, `{}`)
	if len(tb3.fatals) != 1 {
		t.Fatalf("多余内容应 Fatalf：%v", tb3.fatals)
	}
	tb3b := &fakeTB{}
	JSONEqual(tb3b, `{} x`, `{}`)
	if len(tb3b.fatals) != 1 {
		t.Fatalf("多余非法内容应 Fatalf：%v", tb3b.fatals)
	}
	tb4 := &fakeTB{}
	JSONEqual(tb4, 42, `{}`)
	if len(tb4.fatals) != 1 {
		t.Fatalf("不支持类型应 Fatalf：%v", tb4.fatals)
	}
}

func TestFormatJSON(t *testing.T) {
	if got := formatJSON(`{ "a" : 1 }`); got != `{"a":1}` {
		t.Fatalf("JSON 压缩不匹配：%q", got)
	}
	if got := formatJSON(42); got != "42" {
		t.Fatalf("非 JSON 原样输出不匹配：%q", got)
	}
	if got := formatJSON("{bad}"); got != "{bad}" {
		t.Fatalf("非法 JSON 原样输出不匹配：%q", got)
	}
}

func TestJSONNumbersEqualInvalid(t *testing.T) {
	if !jsonNumbersEqual("bad", "bad") {
		t.Fatal("非法数字应退化为字面比较")
	}
	if jsonNumbersEqual("bad", "worse") {
		t.Fatal("非法数字字面不同不应相等")
	}
}
