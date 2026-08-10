package core

import (
	"errors"
	"fmt"
	"strings"
	"testing"

	"github.com/lcylpzls/errx"
)

// checkRequirePasses 验证 Require 断言通过时不产生任何失败。
func checkRequirePasses(t *testing.T, fn func(*fakeTB)) {
	t.Helper()
	tb := &fakeTB{}
	fn(tb)
	if tb.failed() {
		t.Fatalf("断言应通过：errors=%v fatals=%v", tb.errors, tb.fatals)
	}
}

// checkRequireFails 验证 Require 断言失败时仅调用一次 Fatalf。
func checkRequireFails(t *testing.T, fn func(*fakeTB)) {
	t.Helper()
	tb := &fakeTB{}
	fn(tb)
	if len(tb.fatals) != 1 || len(tb.errors) != 0 {
		t.Fatalf("断言失败应仅调用一次 Fatalf：errors=%v fatals=%v", tb.errors, tb.fatals)
	}
}

func TestRequireEqual(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireEqual(tb, 1, 1) })
	checkRequireFails(t, func(tb *fakeTB) { RequireEqual(tb, 1, 2) })
}

func TestRequireNotEqual(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireNotEqual(tb, 1, 2) })
	checkRequireFails(t, func(tb *fakeTB) { RequireNotEqual(tb, 1, 1) })
}

func TestRequireTrueFalse(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireTrue(tb, true) })
	checkRequireFails(t, func(tb *fakeTB) { RequireTrue(tb, false) })
	checkRequirePasses(t, func(tb *fakeTB) { RequireFalse(tb, false) })
	checkRequireFails(t, func(tb *fakeTB) { RequireFalse(tb, true) })
}

func TestRequireNilNotNil(t *testing.T) {
	var p *int
	checkRequirePasses(t, func(tb *fakeTB) { RequireNil(tb, nil) })
	checkRequirePasses(t, func(tb *fakeTB) { RequireNil(tb, p) })
	checkRequireFails(t, func(tb *fakeTB) { RequireNil(tb, 42) })
	checkRequirePasses(t, func(tb *fakeTB) { RequireNotNil(tb, 42) })
	checkRequireFails(t, func(tb *fakeTB) { RequireNotNil(tb, p) })
}

func TestRequireErrorNoError(t *testing.T) {
	sentinel := errors.New("失败")
	checkRequirePasses(t, func(tb *fakeTB) { RequireError(tb, sentinel) })
	checkRequireFails(t, func(tb *fakeTB) { RequireError(tb, nil) })
	checkRequirePasses(t, func(tb *fakeTB) { RequireNoError(tb, nil) })
	checkRequireFails(t, func(tb *fakeTB) { RequireNoError(tb, sentinel) })
}

func TestRequireErrorIs(t *testing.T) {
	sentinel := errors.New("哨兵")
	wrapped := fmt.Errorf("包装：%w", sentinel)
	checkRequirePasses(t, func(tb *fakeTB) { RequireErrorIs(tb, wrapped, sentinel) })
	checkRequireFails(t, func(tb *fakeTB) { RequireErrorIs(tb, errors.New("其他"), sentinel) })
}

func TestRequireEmptyNotEmpty(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireEmpty(tb, "") })
	checkRequireFails(t, func(tb *fakeTB) { RequireEmpty(tb, "x") })
	checkRequirePasses(t, func(tb *fakeTB) { RequireNotEmpty(tb, "x") })
	checkRequireFails(t, func(tb *fakeTB) { RequireNotEmpty(tb, "") })
}

func TestRequireLen(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireLen(tb, []int{1, 2}, 2) })
	checkRequireFails(t, func(tb *fakeTB) { RequireLen(tb, []int{1, 2}, 3) })
}

func TestRequireErrCodeErrKindErrFields(t *testing.T) {
	e := errx.New(errx.KindInvalid, "TEST", "测试错误")
	checkRequirePasses(t, func(tb *fakeTB) { RequireErrCode(tb, e, "TEST") })
	checkRequireFails(t, func(tb *fakeTB) { RequireErrCode(tb, e, "OTHER") })
	checkRequirePasses(t, func(tb *fakeTB) { RequireErrKind(tb, e, errx.KindInvalid) })
	checkRequireFails(t, func(tb *fakeTB) { RequireErrKind(tb, e, errx.KindInternal) })
	ef := e.WithField("k", "v")
	checkRequirePasses(t, func(tb *fakeTB) {
		RequireErrFields(tb, ef, errx.KV{Key: "k", Value: "v"})
	})
	checkRequireFails(t, func(tb *fakeTB) {
		RequireErrFields(tb, ef, errx.KV{Key: "k", Value: "x"})
	})
}

func TestRequireJSONEqual(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) {
		RequireJSONEqual(tb, `{"a":1}`, `{"a":1}`)
	})
	checkRequireFails(t, func(tb *fakeTB) {
		RequireJSONEqual(tb, `{"a":1}`, `{"a":2}`)
	})
}

func TestRequireContains(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireContains(tb, "abcdef", "bcd") })
	checkRequireFails(t, func(tb *fakeTB) { RequireContains(tb, "abcdef", "zzz") })
}

func TestRequireNotContains(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireNotContains(tb, "abc", "z") })
	checkRequireFails(t, func(tb *fakeTB) { RequireNotContains(tb, "abc", "b") })
}

func TestRequireSubset(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireSubset(tb, []int{1, 2, 3}, []int{1, 3}) })
	checkRequireFails(t, func(tb *fakeTB) { RequireSubset(tb, []int{1, 2}, []int{3}) })
}

func TestRequireElementsMatch(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) {
		RequireElementsMatch(tb, []int{1, 2}, []int{2, 1})
	})
	checkRequireFails(t, func(tb *fakeTB) {
		RequireElementsMatch(tb, []int{1}, []int{2})
	})
}

func TestRequirePanics(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequirePanics(tb, func() { panic("x") }) })
	checkRequireFails(t, func(tb *fakeTB) { RequirePanics(tb, func() {}) })
}

func TestRequirePanicsWithValue(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) {
		RequirePanicsWithValue(tb, "x", func() { panic("x") })
	})
	checkRequireFails(t, func(tb *fakeTB) {
		RequirePanicsWithValue(tb, "x", func() { panic("y") })
	})
}

func TestRequireNotPanics(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireNotPanics(tb, func() {}) })
	checkRequireFails(t, func(tb *fakeTB) { RequireNotPanics(tb, func() { panic("x") }) })
}

func TestRequireApprox(t *testing.T) {
	checkRequirePasses(t, func(tb *fakeTB) { RequireApprox(tb, 1.0, 1.05, 0.1) })
	checkRequireFails(t, func(tb *fakeTB) { RequireApprox(tb, 1.0, 2.0, 0.1) })
}

func TestRequireFatalTB(t *testing.T) {
	tb := &fakeTB{}
	fatalTB{tb}.Errorf("格式 %d", 1)
	if len(tb.fatals) != 1 || !strings.Contains(tb.fatals[0], "格式 1") {
		t.Fatalf("Errorf 应转发为 Fatalf：%v", tb.fatals)
	}
}
