# API 快照

> 随版本更新。v0.6.0 快照如下；新版本发布后同步替换。

## v0.6.0

### TB 接口

```go
type TB interface {
    Helper()
    Errorf(format string, args ...any)
    Fatalf(format string, args ...any)
    Cleanup(func())
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

### errx 断言

```go
func ErrCode(t TB, err error, code errx.Code)
func ErrKind(t TB, err error, kind errx.Kind)
func ErrFields(t TB, err error, kvs ...errx.KV)
```

- `ErrCode` 沿错误链匹配错误码（errx.Is）；
- `ErrKind` 取错误链第一个结构化错误的分类（errx.KindOf）；
- `ErrFields` 要求结构化错误包含全部指定键值对（顺序无关）。

### JSON 语义相等

```go
func JSONEqual(t TB, got, want any)
```

- 接受 string / []byte；忽略键序与空白；
- 数字按数值精确比较（`1` 与 `1.0`、`1e5` 与 `100000` 相等）；
- 非法 JSON、多余内容或不支持的类型视为断言误用（Fatalf）。

### 集合/字符串断言

```go
func Contains(t TB, container, elem any)
func NotContains(t TB, container, elem any)
func Subset(t TB, list, sublist any)
func ElementsMatch(t TB, listA, listB any)
```

- `Contains`：字符串检查子串，切片/数组检查元素，map 检查键；
- `Subset`：list 包含 sublist 全部元素（多重性一致，顺序无关）；
- `ElementsMatch`：两个集合元素一致（顺序无关、多重性一致）；
- 不支持的容器类型视为断言误用（Fatalf）。

### 表格驱动、恐慌与数值近似

```go
func RunCases[T any](t *testing.T, name string, cases []T, fn func(t *testing.T, c T))
func Panics(t TB, fn func())
func PanicsWithValue(t TB, want any, fn func())
func NotPanics(t TB, fn func())
func Approx(t TB, got, want, tolerance float64)
```

- `RunCases` 为每个用例生成 `name[序号]` 子测试；
- `Panics` / `PanicsWithValue` / `NotPanics` 基于 recover；
- `Approx` 断言 `|got-want| <= tolerance`，负容差视为误用。

### 辅助

```go
func CaptureStdout(fn func()) string
func CaptureStderr(fn func()) string
func TempEnv(t TB, envs ...string)
func Concurrently(t TB, n int, fn func())
```

- `CaptureStdout` / `CaptureStderr`：临时替换输出流，执行后恢复；
  仅捕获使用 os.Stdout/os.Stderr 全局变量的输出，不可用于并行测试；
- `TempEnv`：成对 KEY/VALUE 设置环境变量，测试结束时恢复原值
  （不存在的键恢复为未设置）；
- `Concurrently`：并发运行 fn 共 n 次并等待完成，配合 `-race` 检测竞态。

### 语义

- `Equal` 使用 `reflect.DeepEqual`；失败输出实际/期望值与类型；
- `Nil` 认 nil 接口、nil 指针/切片/map/函数/通道；
- `Empty` 认零值（空字符串、零数值、nil 指针、空容器）；
- `Len` 支持 string/array/slice/map/chan；
- 失败消息全部使用简体中文。
