# 路线图

## 版本规划

| 版本 | 内容 | 状态 |
| --- | --- | --- |
| v0.1.0 | 核心断言：Equal/NotEqual/True/False/Nil/NotNil/Error/NoError/ErrorIs/Empty/NotEmpty/Len | 进行中 |
| v0.2.0 | errx 断言（ErrCode/ErrKind/ErrFields）与 JSON 语义相等 | 计划 |
| v0.3.0 | 集合/字符串断言：Contains/NotContains/Subset/ElementsMatch | 计划 |
| v0.4.0 | 表格驱动 RunCases、Panics/NotPanics、数值近似 Approx | 计划 |
| v0.5.0 | 捕获/临时环境/并发辅助与基准 | 计划 |
| v0.6.0+ | 终审、自我检查、打磨至 1.0 候选 | 计划 |

## 完成定义（Definition of Done）

- 根包语句覆盖率 100%，测试乱序（`-shuffle=on`）通过；
- race / vet / staticcheck / govulncheck 全绿；
- fuzz 目标短跑 5s 通过（如版本内有目标）；
- 三平台 CI + Linux 多发行版容器矩阵 + 示例模块全绿；
- 文档、CHANGELOG 同步更新；
- 本地全绿后提交并打 tag，由 Release 工作流发布。

## 1.0 候选标准

- 路线图全部完成；
- API 冻结评审通过（签名/语义/失败消息/文档快照）；
- 自检清单（设计/测试/文档/性能/安全）逐项确认；
- 达到后停止自动推进，**v1.0.0 是否发布由维护者决定**。
