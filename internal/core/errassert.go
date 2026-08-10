package core

import (
	"reflect"

	"github.com/lcylpzls/errx"
)

// ErrCode 断言错误链中存在指定 errx 错误码。
func ErrCode(t TB, err error, code errx.Code) {
	t.Helper()
	if !errx.Is(err, code) {
		t.Errorf("期望错误码 %s，实际 %s", code, formatValue(err))
	}
}

// ErrKind 断言错误链中第一个结构化错误的分类为 kind。
func ErrKind(t TB, err error, kind errx.Kind) {
	t.Helper()
	if got := errx.KindOf(err); got != kind {
		t.Errorf("期望错误分类 %s，实际 %s（%s）", kind, got, formatValue(err))
	}
}

// ErrFields 断言结构化错误包含全部指定字段（键值对，顺序无关）。
func ErrFields(t TB, err error, kvs ...errx.KV) {
	t.Helper()
	e, ok := errx.As(err)
	if !ok {
		t.Errorf("期望错误链包含 errx 结构化错误，实际 %s", formatValue(err))
		return
	}
	fields := e.Fields()
	for _, kv := range kvs {
		if !hasField(fields, kv) {
			t.Errorf("期望错误字段 %s=%v，实际字段 %v", kv.Key, kv.Value, fields)
		}
	}
}

// hasField 判断字段列表是否包含指定键值对。
func hasField(fields []errx.KV, want errx.KV) bool {
	for _, f := range fields {
		if f.Key == want.Key && reflect.DeepEqual(f.Value, want.Value) {
			return true
		}
	}
	return false
}
