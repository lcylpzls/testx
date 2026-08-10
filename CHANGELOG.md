# 更新日志

## [v1.4.4] - 2026-08-10

### 变更

- 家族正式基线锁定：依赖统一指向 v1 基线已发布版本（errx v1.5.5 / logx v1.3.2 / testx v1.4.3 / validx v1.2.4 / cryptox v1.0.2 / confx v1.0.2 / webx v1.5.4 等），此后家族依赖不再前进。

### 质量

- 全部库包语句覆盖率保持 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v1.4.3] - 2026-08-10

### 变更

- 家族依赖最终对齐到 v1 正式版基线（errx v1.5.4 / logx v1.3.1 / testx v1.4.2 / validx v1.2.3 / confx v1.0.1 / cryptox v1.0.1 等），无 API 变更。

### 质量

- 全部库包语句覆盖率保持 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v1.4.2] - 2026-08-10

### 变更

- 家族依赖统一对齐到最新基线（errx v1.5.4 / logx v1.3.0 / testx v1.4.1 / validx v1.2.2 等），无 API 变更。

### 质量

- 全部库包语句覆盖率保持 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v1.4.1] - 2026-08-10

### 变更

- JSON 解析辅助错误统一 errx 化（新增 CodeInvalidJSON），对外错误带结构化 code/kind，消息保持原语义。

### 质量

- 覆盖率维持基线；race / vet / staticcheck / govulncheck 全绿。

## [v1.4.0] - 2026-08-10

### 新增

- 数值比较断言：`Greater` / `GreaterOrEqual` / `Less` /
  `LessOrEqual` 及对应 `Require*` 版本（支持全部整数、无符号与浮点类型，
  非数值输入按失败处理）。

### 质量

- race / vet / staticcheck / govulncheck 全绿；覆盖率维持测试基座自身基线。

## [v1.3.0] - 2026-08-10

### 新增

- `testx/logx` 可选子包：日志捕获器 `New` / `Snapshot`（实现完整
  logx.Logger）与断言 `AssertLogged` / `AssertLoggedContains`，
  供家族各库测试日志行为。

### 质量

- 根包与子包覆盖率均 100%；race / vet / staticcheck / fuzz /
  govulncheck 全绿。

## [v1.2.4] - 2026-08-10

### 变更

- go 指令与 CI/Release 工作流统一为 Go 1.26.5；
- README Go 版本徽章同步更新。

## [v1.2.3] - 2026-08-10

### 变更

- go 指令与 CI/Release 工作流统一为精确版本 `1.21.0`。

## [v1.2.2] - 2026-08-10

### 变更

- 家族统一 Go 1.21：全部 go.mod 与 CI/Release 工作流版本号对齐 1.21。

## [v1.2.1] - 2026-08-10

### 修复

- go.mod 的 go 指令降回 `go 1.21`：testx 本身不依赖新语言特性，
  避免家族内 Go 1.21 矩阵（confx / errx 等）因依赖门槛失败。

## [v1.2.0] - 2026-08-10

### 新增

- `Require*` 致命断言变体：所有断言均提供 Require 版本（失败时
  立即 `Fatalf` 终止测试），与既有非致命断言并存；
- 提供 `RequireEqual/RequireNoError/RequireErrCode` 等 24 个变体，
  覆盖基础断言、errx 断言、JSON、集合、恐慌与数值近似。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

## [v1.1.0] - 2026-08-10

### 变更

- 家族依赖同步：`errx` v1.3.2 → v1.4.0（纯增量，新增可选全局
  MetricsHook，不影响现有断言 API）；
- 示例模块间接依赖同步升级。

### 质量

- 根包语句覆盖率 100%；race / vet / staticcheck / fuzz / govulncheck 全绿。

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
