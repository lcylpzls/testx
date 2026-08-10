package core

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math/big"
	"reflect"

	"github.com/lcylpzls/errx"
)

// CodeInvalidJSON JSON 解析失败（解码/多余内容）。
const CodeInvalidJSON errx.Code = "TESTX_INVALID_JSON"

func init() {
	errx.RegisterCode(CodeInvalidJSON, "JSON 解析失败")
	errx.RegisterCodeKind(CodeInvalidJSON, errx.KindInvalid)
}

// JSONEqual 断言两个 JSON（string/[]byte）语义相等：
// 忽略键序与空白，数字按数值比较（1 与 1.0 相等）。
// 输入不是合法 JSON 时视为断言误用，立即终止测试。
func JSONEqual(t TB, got, want any) {
	t.Helper()
	gv, err := decodeJSON(got)
	if err != nil {
		t.Fatalf("实际值不是合法 JSON：%v（%s）", err, formatValue(got))
		return
	}
	wv, err := decodeJSON(want)
	if err != nil {
		t.Fatalf("期望值不是合法 JSON：%v（%s）", err, formatValue(want))
		return
	}
	if !jsonDeepEqual(gv, wv) {
		t.Errorf("期望 JSON 语义相等：\n  实际: %s\n  期望: %s",
			formatJSON(got), formatJSON(want))
	}
}

// decodeJSON 将 string/[]byte 解码为 UseNumber 的通用值，并拒绝多余内容。
func decodeJSON(v any) (any, error) {
	data, ok := jsonData(v)
	if !ok {
		return nil, errx.NewCodef(CodeInvalidJSON, "仅支持 string/[]byte，当前类型 %T", v)
	}
	dec := json.NewDecoder(bytes.NewReader(data))
	dec.UseNumber()
	var out any
	if err := dec.Decode(&out); err != nil {
		return nil, err
	}
	var extra any
	if err := dec.Decode(&extra); !errors.Is(err, io.EOF) {
		if err == nil {
			return nil, errx.NewCode(CodeInvalidJSON, "JSON 后存在多余内容")
		}
		return nil, err
	}
	return out, nil
}

// jsonData 提取 string/[]byte 形式的 JSON 数据。
func jsonData(v any) ([]byte, bool) {
	switch val := v.(type) {
	case string:
		return []byte(val), true
	case []byte:
		return val, true
	default:
		return nil, false
	}
}

// jsonDeepEqual 递归比较 JSON 值：map 无序、数字按数值。
func jsonDeepEqual(a, b any) bool {
	switch av := a.(type) {
	case json.Number:
		bv, ok := b.(json.Number)
		if !ok {
			return false
		}
		return jsonNumbersEqual(av, bv)
	case map[string]any:
		bm, ok := b.(map[string]any)
		if !ok || len(av) != len(bm) {
			return false
		}
		for k, v := range av {
			bv, ok := bm[k]
			if !ok || !jsonDeepEqual(v, bv) {
				return false
			}
		}
		return true
	case []any:
		bs, ok := b.([]any)
		if !ok || len(av) != len(bs) {
			return false
		}
		for i := range av {
			if !jsonDeepEqual(av[i], bs[i]) {
				return false
			}
		}
		return true
	default:
		return reflect.DeepEqual(a, b)
	}
}

// jsonNumbersEqual 使用有理数精确比较 JSON 数字。
func jsonNumbersEqual(a, b json.Number) bool {
	ra, okA := new(big.Rat).SetString(a.String())
	rb, okB := new(big.Rat).SetString(b.String())
	if !okA || !okB {
		// 仅防御非法输入；JSON 解码器产出的数字均可被 Rat 精确解析。
		return a.String() == b.String()
	}
	return ra.Cmp(rb) == 0
}

// formatJSON 压缩 JSON 输出用于失败消息；非 JSON 输入原样格式化。
func formatJSON(v any) string {
	data, ok := jsonData(v)
	if !ok {
		return fmt.Sprint(v)
	}
	var buf bytes.Buffer
	if err := json.Compact(&buf, data); err != nil {
		return fmt.Sprint(v)
	}
	return buf.String()
}
