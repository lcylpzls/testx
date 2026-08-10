package testx

import (
	"github.com/lcylpzls/errx"
	"github.com/lcylpzls/testx/internal/core"
)

type TB = core.TB

const CodeInvalidJSON = core.CodeInvalidJSON

func ErrCode(t TB, err error, code errx.Code)          { core.ErrCode(t, err, code) }
func ErrKind(t TB, err error, kind errx.Kind)          { core.ErrKind(t, err, kind) }
func ErrFields(t TB, err error, kvs ...errx.KV)        { core.ErrFields(t, err, kvs...) }
func CaptureStdout(fn func()) string                   { return core.CaptureStdout(fn) }
func CaptureStderr(fn func()) string                   { return core.CaptureStderr(fn) }
func TempEnv(t TB, envs ...string)                     { core.TempEnv(t, envs...) }
func Concurrently(t TB, n int, fn func())              { core.Concurrently(t, n, fn) }
func Equal(t TB, got, want any)                        { core.Equal(t, got, want) }
func NotEqual(t TB, got, want any)                     { core.NotEqual(t, got, want) }
func True(t TB, v bool)                                { core.True(t, v) }
func False(t TB, v bool)                               { core.False(t, v) }
func Nil(t TB, v any)                                  { core.Nil(t, v) }
func NotNil(t TB, v any)                               { core.NotNil(t, v) }
func Error(t TB, err error)                            { core.Error(t, err) }
func NoError(t TB, err error)                          { core.NoError(t, err) }
func ErrorIs(t TB, err, target error)                  { core.ErrorIs(t, err, target) }
func Empty(t TB, v any)                                { core.Empty(t, v) }
func NotEmpty(t TB, v any)                             { core.NotEmpty(t, v) }
func Len(t TB, v any, want int)                        { core.Len(t, v, want) }
func Greater(t TB, got, want any)                      { core.Greater(t, got, want) }
func GreaterOrEqual(t TB, got, want any)               { core.GreaterOrEqual(t, got, want) }
func Less(t TB, got, want any)                         { core.Less(t, got, want) }
func LessOrEqual(t TB, got, want any)                  { core.LessOrEqual(t, got, want) }
func Contains(t TB, container, elem any)               { core.Contains(t, container, elem) }
func NotContains(t TB, container, elem any)            { core.NotContains(t, container, elem) }
func Subset(t TB, list, sublist any)                   { core.Subset(t, list, sublist) }
func ElementsMatch(t TB, listA, listB any)             { core.ElementsMatch(t, listA, listB) }
func Panics(t TB, fn func())                           { core.Panics(t, fn) }
func PanicsWithValue(t TB, want any, fn func())        { core.PanicsWithValue(t, want, fn) }
func NotPanics(t TB, fn func())                        { core.NotPanics(t, fn) }
func Approx(t TB, got, want, tolerance float64)        { core.Approx(t, got, want, tolerance) }
func JSONEqual(t TB, got, want any)                    { core.JSONEqual(t, got, want) }
func RequireEqual(t TB, got, want any)                 { core.RequireEqual(t, got, want) }
func RequireNotEqual(t TB, got, want any)              { core.RequireNotEqual(t, got, want) }
func RequireTrue(t TB, v bool)                         { core.RequireTrue(t, v) }
func RequireFalse(t TB, v bool)                        { core.RequireFalse(t, v) }
func RequireNil(t TB, v any)                           { core.RequireNil(t, v) }
func RequireNotNil(t TB, v any)                        { core.RequireNotNil(t, v) }
func RequireError(t TB, err error)                     { core.RequireError(t, err) }
func RequireNoError(t TB, err error)                   { core.RequireNoError(t, err) }
func RequireErrorIs(t TB, err, target error)           { core.RequireErrorIs(t, err, target) }
func RequireEmpty(t TB, v any)                         { core.RequireEmpty(t, v) }
func RequireNotEmpty(t TB, v any)                      { core.RequireNotEmpty(t, v) }
func RequireLen(t TB, v any, want int)                 { core.RequireLen(t, v, want) }
func RequireErrCode(t TB, err error, code errx.Code)   { core.RequireErrCode(t, err, code) }
func RequireErrKind(t TB, err error, kind errx.Kind)   { core.RequireErrKind(t, err, kind) }
func RequireErrFields(t TB, err error, kvs ...errx.KV) { core.RequireErrFields(t, err, kvs...) }
func RequireJSONEqual(t TB, got, want any)             { core.RequireJSONEqual(t, got, want) }
func RequireContains(t TB, container, elem any)        { core.RequireContains(t, container, elem) }
func RequireNotContains(t TB, container, elem any)     { core.RequireNotContains(t, container, elem) }
func RequireSubset(t TB, list, sublist any)            { core.RequireSubset(t, list, sublist) }
func RequireElementsMatch(t TB, listA, listB any)      { core.RequireElementsMatch(t, listA, listB) }
func RequirePanics(t TB, fn func())                    { core.RequirePanics(t, fn) }
func RequirePanicsWithValue(t TB, want any, fn func()) { core.RequirePanicsWithValue(t, want, fn) }
func RequireNotPanics(t TB, fn func())                 { core.RequireNotPanics(t, fn) }
func RequireApprox(t TB, got, want, tolerance float64) { core.RequireApprox(t, got, want, tolerance) }
func RequireGreater(t TB, got, want any)               { core.RequireGreater(t, got, want) }
func RequireGreaterOrEqual(t TB, got, want any)        { core.RequireGreaterOrEqual(t, got, want) }
func RequireLess(t TB, got, want any)                  { core.RequireLess(t, got, want) }
func RequireLessOrEqual(t TB, got, want any)           { core.RequireLessOrEqual(t, got, want) }
