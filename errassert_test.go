package testx

import (
	"errors"
	"strings"
	"testing"

	"github.com/lcylpzls/errx"
)

func TestErrCode(t *testing.T) {
	err := errx.NewCode("TEST_CODE", "测试失败")
	ErrCode(&fakeTB{}, err, "TEST_CODE")
	tb := &fakeTB{}
	ErrCode(tb, err, "OTHER_CODE")
	if !tb.failed() {
		t.Fatal("错误码不匹配应失败")
	}
	if !strings.Contains(tb.errors[0], "TEST_CODE") {
		t.Fatalf("失败消息应包含实际错误码：%s", tb.errors[0])
	}
	ErrCode(&fakeTB{}, nil, "TEST_CODE")
}

func TestErrKind(t *testing.T) {
	err := errx.New(errx.KindNotFound, "TEST_CODE", "未找到")
	ErrKind(&fakeTB{}, err, errx.KindNotFound)
	tb := &fakeTB{}
	ErrKind(tb, err, errx.KindInvalid)
	if !tb.failed() {
		t.Fatal("分类不匹配应失败")
	}
	if !strings.Contains(tb.errors[0], "not_found") {
		t.Fatalf("失败消息应包含实际分类：%s", tb.errors[0])
	}
	ErrKind(&fakeTB{}, nil, errx.KindUnknown)
}

func TestErrFields(t *testing.T) {
	err := errx.NewCode("TEST_CODE", "失败").
		WithField("order_id", "10086").
		WithField("retry", 3)
	ErrFields(&fakeTB{}, err,
		errx.KV{Key: "order_id", Value: "10086"},
		errx.KV{Key: "retry", Value: 3},
	)
	tb := &fakeTB{}
	ErrFields(tb, err, errx.KV{Key: "missing", Value: "x"})
	if !tb.failed() {
		t.Fatal("缺失字段应失败")
	}
	tb2 := &fakeTB{}
	ErrFields(tb2, errors.New("普通错误"), errx.KV{Key: "a", Value: 1})
	if !tb2.failed() {
		t.Fatal("非 errx 错误应失败")
	}
}

func TestHasField(t *testing.T) {
	fields := []errx.KV{{Key: "a", Value: 1}}
	if !hasField(fields, errx.KV{Key: "a", Value: 1}) {
		t.Fatal("应命中键值对")
	}
	if hasField(fields, errx.KV{Key: "a", Value: 2}) {
		t.Fatal("值不同不应命中")
	}
	if hasField(fields, errx.KV{Key: "b", Value: 1}) {
		t.Fatal("键不同不应命中")
	}
}
