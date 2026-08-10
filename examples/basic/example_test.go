package basic_test

import (
	"errors"
	"fmt"
	"testing"

	"github.com/lcylpzls/errx"
	"github.com/lcylpzls/testx"
)

func TestExample(t *testing.T) {
	testx.Equal(t, 1+1, 2)
	testx.NotNil(t, "有值")
	testx.NoError(t, nil)

	err := errors.New("失败")
	testx.Error(t, err)
	testx.ErrorIs(t, err, err)
	testx.Len(t, []int{1, 2, 3}, 3)

	e := errx.NewCode("EXAMPLE_FAIL", "示例失败").WithField("id", 1)
	testx.ErrCode(t, e, "EXAMPLE_FAIL")
	testx.ErrFields(t, e, errx.KV{Key: "id", Value: 1})
	testx.JSONEqual(t, `{"a":1}`, `{"a":1.0}`)
	testx.Contains(t, "你好世界", "世界")
	testx.Subset(t, []int{1, 2, 3}, []int{3, 1})
	testx.ElementsMatch(t, []string{"a", "b"}, []string{"b", "a"})
	testx.Panics(t, func() { panic("boom") })
	testx.Approx(t, 1.0, 1.01, 0.1)

	testx.RunCases(t, "加法", []struct {
		a, b, want int
	}{
		{1, 2, 3},
		{2, 3, 5},
	}, func(t *testing.T, c struct{ a, b, want int }) {
		testx.Equal(t, c.a+c.b, c.want)
	})

	got := testx.CaptureStdout(func() {
		fmt.Print("捕获")
	})
	testx.Equal(t, got, "捕获")
	testx.TempEnv(t, "EXAMPLE_ENV", "值")
	testx.Concurrently(t, 4, func() {})
}
