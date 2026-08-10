package core

import "github.com/lcylpzls/errx"

// fatalTB 将断言失败从 Errorf 升级为 Fatalf，供 Require 系列断言使用。
type fatalTB struct {
	TB
}

// Errorf 以 Fatalf 形式记录失败并立即终止测试。
func (f fatalTB) Errorf(format string, args ...any) {
	f.Fatalf(format, args...)
}

// RequireEqual 断言 got 与 want 深度相等，失败时立即终止测试。
func RequireEqual(t TB, got, want any) { Equal(fatalTB{t}, got, want) }

// RequireNotEqual 断言 got 与 want 不相等，失败时立即终止测试。
func RequireNotEqual(t TB, got, want any) { NotEqual(fatalTB{t}, got, want) }

// RequireTrue 断言值为 true，失败时立即终止测试。
func RequireTrue(t TB, v bool) { True(fatalTB{t}, v) }

// RequireFalse 断言值为 false，失败时立即终止测试。
func RequireFalse(t TB, v bool) { False(fatalTB{t}, v) }

// RequireNil 断言值为 nil，失败时立即终止测试。
func RequireNil(t TB, v any) { Nil(fatalTB{t}, v) }

// RequireNotNil 断言值非 nil，失败时立即终止测试。
func RequireNotNil(t TB, v any) { NotNil(fatalTB{t}, v) }

// RequireError 断言 err 非 nil，失败时立即终止测试。
func RequireError(t TB, err error) { Error(fatalTB{t}, err) }

// RequireNoError 断言 err 为 nil，失败时立即终止测试。
func RequireNoError(t TB, err error) { NoError(fatalTB{t}, err) }

// RequireErrorIs 断言错误链中包含 target，失败时立即终止测试。
func RequireErrorIs(t TB, err, target error) { ErrorIs(fatalTB{t}, err, target) }

// RequireEmpty 断言值为零值/空值，失败时立即终止测试。
func RequireEmpty(t TB, v any) { Empty(fatalTB{t}, v) }

// RequireNotEmpty 断言值非零值/非空，失败时立即终止测试。
func RequireNotEmpty(t TB, v any) { NotEmpty(fatalTB{t}, v) }

// RequireLen 断言容器长度，失败时立即终止测试。
func RequireLen(t TB, v any, want int) { Len(fatalTB{t}, v, want) }

// RequireErrCode 断言错误链中存在指定 errx 错误码，失败时立即终止测试。
func RequireErrCode(t TB, err error, code errx.Code) { ErrCode(fatalTB{t}, err, code) }

// RequireErrKind 断言错误分类为 kind，失败时立即终止测试。
func RequireErrKind(t TB, err error, kind errx.Kind) { ErrKind(fatalTB{t}, err, kind) }

// RequireErrFields 断言错误包含全部指定字段，失败时立即终止测试。
func RequireErrFields(t TB, err error, kvs ...errx.KV) { ErrFields(fatalTB{t}, err, kvs...) }

// RequireJSONEqual 断言 JSON 语义相等，失败时立即终止测试。
func RequireJSONEqual(t TB, got, want any) { JSONEqual(fatalTB{t}, got, want) }

// RequireContains 断言容器包含元素，失败时立即终止测试。
func RequireContains(t TB, container, elem any) { Contains(fatalTB{t}, container, elem) }

// RequireNotContains 断言容器不包含元素，失败时立即终止测试。
func RequireNotContains(t TB, container, elem any) { NotContains(fatalTB{t}, container, elem) }

// RequireSubset 断言 list 包含 sublist 全部元素，失败时立即终止测试。
func RequireSubset(t TB, list, sublist any) { Subset(fatalTB{t}, list, sublist) }

// RequireElementsMatch 断言两个集合元素一致，失败时立即终止测试。
func RequireElementsMatch(t TB, listA, listB any) { ElementsMatch(fatalTB{t}, listA, listB) }

// RequirePanics 断言 fn 触发 panic，失败时立即终止测试。
func RequirePanics(t TB, fn func()) { Panics(fatalTB{t}, fn) }

// RequirePanicsWithValue 断言 panic 值与 want 相等，失败时立即终止测试。
func RequirePanicsWithValue(t TB, want any, fn func()) {
	PanicsWithValue(fatalTB{t}, want, fn)
}

// RequireNotPanics 断言 fn 不触发 panic，失败时立即终止测试。
func RequireNotPanics(t TB, fn func()) { NotPanics(fatalTB{t}, fn) }

// RequireApprox 断言数值近似相等，失败时立即终止测试。
func RequireApprox(t TB, got, want, tolerance float64) {
	Approx(fatalTB{t}, got, want, tolerance)
}

// RequireGreater 断言 got 大于 want，失败时立即终止测试。
func RequireGreater(t TB, got, want any) { Greater(fatalTB{t}, got, want) }

// RequireGreaterOrEqual 断言 got 大于等于 want，失败时立即终止测试。
func RequireGreaterOrEqual(t TB, got, want any) { GreaterOrEqual(fatalTB{t}, got, want) }

// RequireLess 断言 got 小于 want，失败时立即终止测试。
func RequireLess(t TB, got, want any) { Less(fatalTB{t}, got, want) }

// RequireLessOrEqual 断言 got 小于等于 want，失败时立即终止测试。
func RequireLessOrEqual(t TB, got, want any) { LessOrEqual(fatalTB{t}, got, want) }
