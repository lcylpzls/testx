# testx

轻量、零魔法的 Go 测试断言与辅助库：与 errx 结构化错误天然打通，
消除家族各库重复编写 fake/helper 的问题。

[![Go Version](https://img.shields.io/badge/Go-%3E%3D1.26-00ADD8?style=flat&logo=go)](https://go.dev/)
[![License](https://img.shields.io/badge/license-MIT-green.svg)](LICENSE)

## 快速开始

```go
func TestGreeting(t *testing.T) {
    got := greeting("小明")
    testx.Equal(t, got, "你好，小明！")
    testx.NotNil(t, got)
}
```

## 核心特性

- 核心断言：Equal / NotEqual / True / False / Nil / NotNil / Error /
  NoError / ErrorIs / Empty / NotEmpty / Len；
- errx 结构化断言：错误码 / 分类 / 字段（v0.2.0 起）；
- JSON 语义相等（v0.2.0 起）；
- 集合/字符串断言与表格驱动（v0.3.0 / v0.4.0 起）；
- 输出捕获、临时环境与并发辅助（v0.5.0 起）；
- 失败消息全部简体中文，断言基于可替换的 `TB` 接口。

## 文档

- [docs/research.md](docs/research.md) — 竞品调研与取舍
- [docs/design.md](docs/design.md) — 设计
- [docs/architecture.md](docs/architecture.md) — 架构
- [docs/api.md](docs/api.md) — API 快照
- [docs/roadmap.md](docs/roadmap.md) — 路线图

## License

MIT © [lcylpzls](https://github.com/lcylpzls)
