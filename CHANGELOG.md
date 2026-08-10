# 更新日志

## [v0.2.0] - 2026-08-10

### 新增

- errx 结构化断言：`ErrCode`（错误码链匹配）、`ErrKind`（分类）、
  `ErrFields`（字段包含，顺序无关）；
- JSON 语义相等：`JSONEqual` 忽略键序与空白，数字按数值精确比较
  （`1` 与 `1.0`、`1e5` 与 `100000` 相等），非法输入按误用 Fatalf；
- `FuzzJSONEqual` 接入 CI。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.1.0] - 2026-08-10

### 新增

- `TB` 接口（Helper/Errorf/Fatalf 子集），`*testing.T` 天然满足；
- 核心断言：`Equal` / `NotEqual` / `True` / `False` / `Nil` / `NotNil` /
  `Error` / `NoError` / `ErrorIs` / `Empty` / `NotEmpty` / `Len`；
- 失败消息简体中文，含实际/期望值与类型；
- fuzz 目标（`FuzzFormatValue`）接入 CI；
- 三平台 CI + Linux 多发行版容器矩阵 + Release 工作流。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。
