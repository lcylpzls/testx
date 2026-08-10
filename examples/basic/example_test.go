package basic_test

import (
	"errors"
	"testing"

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
}
