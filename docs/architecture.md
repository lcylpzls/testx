# 架构

## 1. 包内模块

```text
testx（根包）
├── tb.go        TB 接口
├── assert.go    核心断言与值格式化
├── errassert.go errx 结构化断言（v0.2.0）
├── json.go      JSON 语义相等（v0.2.0）
├── collect.go   集合/字符串断言（v0.3.0）
├── table.go     表格驱动与近似/恐慌断言（v0.4.0）
└── capture.go   输出捕获、临时环境与并发辅助（v0.5.0）
```

依赖方向：

```text
assert.go ──→ tb.go
errassert.go ──→ assert.go ──→ tb.go
collect.go / table.go / capture.go ──→ assert.go
```

## 2. 关键设计

- **TB 抽象**：对外只要求 `Helper` / `Errorf` / `Fatalf` 三个方法，
  `*testing.T` 天然满足，库自身测试用替身验证失败路径；
- **值格式化**：`formatValue` 统一输出类型与可读值，字符串带引号，
  结构体/切片输出 `%#v`；
- **零全局状态**：所有辅助（如输出捕获）通过标准库机制与 `t.Cleanup`
  管理，无包级可变配置；
- **并发安全**：断言只读输入；捕获辅助在测试协程内使用，
  `Concurrently` 等待全部协程完成。

## 3. 后续演进扩展点

- errx 断言依赖 `errx.As/Is/CodeOf/KindOf`，失败消息可扩展字段输出；
- JSON 相等基于 `encoding/json` 反序列化后 `reflect.DeepEqual`；
- 表格驱动使用泛型，不依赖反射。
