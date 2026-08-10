# 更新日志

## [v1.0.0] - 2026-08-10

### 发布

- testx 正式发布 1.0.0：API 冻结，破坏性变更必须提升主版本；
- 依赖约定生效：仅依赖 `errx v1.3.x`（1.x），无 0.x 依赖；
- 终审清单全部通过，示例与文档对齐 1.0.0 发布形态。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.6.2] - 2026-08-10

### 文档与约定

- API 快照同步至 v0.6.2；
- 明确依赖约定：仅依赖 `errx v1.3.x`，v1.0.0 起 API 冻结。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.6.1] - 2026-08-10

### 文档

- API 快照同步至 v0.6.0。

## [v0.6.0] - 2026-08-10

### 修复

- 输出捕获改为边写边读：超过管道缓冲区（64 KiB）的输出不再阻塞；
- `panic(nil)` 与未 panic 现在可区分（`PanicsWithValue(nil)` 可通过）；
- `Approx` 对 NaN 实际值/期望值明确失败。

### 改进

- 新增 `docs/final-review.md` 1.0 候选终审清单与 Issue 模板；
- README 增加 CI 徽章。

### 结论

- testx 达到 1.0 候选标准；**v1.0.0 是否发布由维护者决定**。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.5.0] - 2026-08-10

### 新增

- 输出捕获：`CaptureStdout` / `CaptureStderr`（临时替换输出流，
  panic 时也恢复）；
- 临时环境：`TempEnv` 成对设置并在测试结束恢复原值；
- 并发辅助：`Concurrently` 并发运行并等待完成，配合 `-race` 检测竞态；
- TB 接口新增 `Cleanup`；
- 基准测试与 `BENCHMARKS.md`（Equal 通过路径零分配）。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.4.0] - 2026-08-10

### 新增

- 泛型表格驱动 `RunCases`：每个用例生成 `name[序号]` 子测试；
- 恐慌断言：`Panics` / `PanicsWithValue` / `NotPanics`；
- 数值近似：`Approx`（容差非负，负容差按误用 Fatalf）。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v0.3.0] - 2026-08-10

### 新增

- 集合/字符串断言：`Contains` / `NotContains`（字符串子串、切片/数组
  元素、map 键）；
- 集合关系断言：`Subset`（多重性一致、顺序无关）与 `ElementsMatch`
  （顺序无关、多重性一致）；
- 不支持的容器类型按断言误用 Fatalf。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

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
