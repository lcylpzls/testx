# API 快照

> 随版本更新。v0.1.0 快照如下；新版本发布后同步替换。

## v0.1.0

### TB 接口

```go
type TB interface {
    Helper()
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
}
```

`*testing.T` 天然满足；断言失败用 `Errorf`，误用用 `Fatalf`。

### 断言

```go
func Equal(t TB, got, want any)
func NotEqual(t TB, got, want any)
func True(t TB, v bool)
func False(t TB, v bool)
func Nil(t TB, v any)
func NotNil(t TB, v any)
func Error(t TB, err error)
func NoError(t TB, err error)
func ErrorIs(t TB, err, target error)
func Empty(t TB, v any)
func NotEmpty(t TB, v any)
func Len(t TB, v any, want int)
```

### 语义

- `Equal` 使用 `reflect.DeepEqual`；失败输出实际/期望值与类型；
- `Nil` 认 nil 接口、nil 指针/切片/map/函数/通道；
- `Empty` 认零值（空字符串、零数值、nil 指针、空容器）；
- `Len` 支持 string/array/slice/map/chan；
- 失败消息全部使用简体中文。
