# 1.0 候选终审

> 本清单用于确认 testx 达到 1.0 候选标准；**v1.0.0 是否发布由维护者决定**。

## 1. API 冻结

- [x] 公开 API 签名稳定：TB/断言/辅助语义明确；
- [x] 失败消息与误用（Fatalf）约定形成稳定快照；
- [x] pre-1.0 兼容承诺：v0.1.0 起的核心断言行为无意外破坏。

## 2. 质量门禁

- [x] 根包语句覆盖率 100%；
- [x] 测试乱序（`-shuffle=on`）、race 全平台通过；
- [x] vet / staticcheck / govulncheck 通过；
- [x] fuzz 目标短跑 5s 通过；
- [x] 示例模块（含 errx/JSON/集合/表格/捕获/并发）全绿。

## 3. 设计确认

- [x] 零全局可变状态；断言只读输入；
- [x] TB 接口可替换，失败路径可用替身验证；
- [x] 捕获辅助在 panic 与超量输出下均安全；
- [x] 并发辅助配合 `-race` 可检测竞态。

## 4. 性能

- [x] `BENCHMARKS.md` 记录断言基准；
- [x] Equal 通过路径零分配。

## 5. 文档与安全

- [x] README / docs/api.md / docs/roadmap.md 一致；
- [x] SECURITY.md / CONTRIBUTING.md / CODEOWNERS / Issue 模板齐全；
- [x] 发布流程：tag 触发 Release，CI 全绿后发布。

## 结论

testx 已通过 1.0 候选终审清单，达到 1.0 候选标准。
**v1.0.0 是否发布由维护者决定**；确认发布前不再自动推进版本。
