package testx_test

import (
	"errors"
	"testing"

	"github.com/lcylpzls/errx"
	"github.com/lcylpzls/testx"
)

// fakeTB 实现 testx.TB，用于冒烟测试。
type fakeTB struct {
	errors []string
}

func (f *fakeTB) Helper()                           {}
func (f *fakeTB) Errorf(format string, args ...any) { f.errors = append(f.errors, format) }
func (f *fakeTB) Fatalf(format string, args ...any) { f.errors = append(f.errors, format) }
func (f *fakeTB) Cleanup(func())                    {}

// TestPublicAPI 黑盒冒烟测试：覆盖根包全部转发函数、类型别名与常量。
func TestPublicAPI(t *testing.T) {
	tb := &fakeTB{}
	smokeErr := errx.New(errx.KindBusiness, errx.Code("smoke"), "冒烟")

	testx.ErrCode(tb, smokeErr, errx.Code("smoke"))
	testx.ErrKind(tb, smokeErr, errx.KindBusiness)
	testx.ErrFields(tb, smokeErr)
	_ = testx.CaptureStdout(func() {})
	_ = testx.CaptureStderr(func() {})
	testx.TempEnv(tb)
	testx.Concurrently(tb, 2, func() {})
	testx.Equal(tb, 1, 1)
	testx.NotEqual(tb, 1, 2)
	testx.True(tb, true)
	testx.False(tb, false)
	testx.Nil(tb, nil)
	testx.NotNil(tb, 1)
	testx.Error(tb, errors.New("x"))
	testx.NoError(tb, nil)
	testx.ErrorIs(tb, smokeErr, smokeErr)
	testx.Empty(tb, []int{})
	testx.NotEmpty(tb, []int{1})
	testx.Len(tb, []int{1, 2}, 2)
	testx.Greater(tb, 2, 1)
	testx.GreaterOrEqual(tb, 2, 2)
	testx.Less(tb, 1, 2)
	testx.LessOrEqual(tb, 2, 2)
	testx.Contains(tb, []int{1, 2}, 1)
	testx.NotContains(tb, []int{1, 2}, 3)
	testx.Subset(tb, []int{1, 2, 3}, []int{1, 2})
	testx.ElementsMatch(tb, []int{1, 2}, []int{2, 1})
	testx.Panics(tb, func() { panic("x") })
	testx.PanicsWithValue(tb, "x", func() { panic("x") })
	testx.NotPanics(tb, func() {})
	testx.Approx(tb, 1.0, 1.01, 0.1)
	testx.JSONEqual(tb, map[string]int{"a": 1}, map[string]int{"a": 1})

	testx.RequireEqual(tb, 1, 1)
	testx.RequireNotEqual(tb, 1, 2)
	testx.RequireTrue(tb, true)
	testx.RequireFalse(tb, false)
	testx.RequireNil(tb, nil)
	testx.RequireNotNil(tb, 1)
	testx.RequireError(tb, errors.New("x"))
	testx.RequireNoError(tb, nil)
	testx.RequireErrorIs(tb, smokeErr, smokeErr)
	testx.RequireEmpty(tb, []int{})
	testx.RequireNotEmpty(tb, []int{1})
	testx.RequireLen(tb, []int{1, 2}, 2)
	testx.RequireErrCode(tb, smokeErr, errx.Code("smoke"))
	testx.RequireErrKind(tb, smokeErr, errx.KindBusiness)
	testx.RequireErrFields(tb, smokeErr)
	testx.RequireJSONEqual(tb, map[string]int{"a": 1}, map[string]int{"a": 1})
	testx.RequireContains(tb, []int{1, 2}, 1)
	testx.RequireNotContains(tb, []int{1, 2}, 3)
	testx.RequireSubset(tb, []int{1, 2, 3}, []int{1, 2})
	testx.RequireElementsMatch(tb, []int{1, 2}, []int{2, 1})
	testx.RequirePanics(tb, func() { panic("x") })
	testx.RequirePanicsWithValue(tb, "x", func() { panic("x") })
	testx.RequireNotPanics(tb, func() {})
	testx.RequireApprox(tb, 1.0, 1.01, 0.1)
	testx.RequireGreater(tb, 2, 1)
	testx.RequireGreaterOrEqual(tb, 2, 2)
	testx.RequireLess(tb, 1, 2)
	testx.RequireLessOrEqual(tb, 2, 2)

	var _ testx.TB = tb
	_ = testx.CodeInvalidJSON
}
