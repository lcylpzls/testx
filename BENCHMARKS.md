# 基准测试

> 采集环境：Windows / AMD Ryzen 5 7600 / Go 1.26.5
> 采集日期：2026-08-10
> 命令：`go test -bench=. -benchmem -run '^$' .`

## BenchmarkEqual

字符串深度相等断言（全部通过路径）：

| 指标 | 数值 |
| --- | --- |
| 耗时 | 15.7 ns/op |
| 内存 | 0 B/op |
| 分配 | 0 allocs/op |

## BenchmarkJSONEqual

JSON 语义相等（键序不同 + 数字 `1`/`1.0`）：

| 指标 | 数值 |
| --- | --- |
| 耗时 | 3680 ns/op |
| 内存 | 3658 B/op |
| 分配 | 80 allocs/op |

## BenchmarkContains

8 元素整数切片包含断言：

| 指标 | 数值 |
| --- | --- |
| 耗时 | 265 ns/op |
| 内存 | 88 B/op |
| 分配 | 9 allocs/op |

## 说明

- 基准仅反映本机相对量级；CI 不设硬性性能门槛；
- 断言库位于测试路径，性能不是首要目标，但通过路径保持零分配
  （Equal）与可预期开销。
