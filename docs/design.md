# 设计

## 1. 定位与范围

testx 是家族生态的测试断言与辅助库：

- 覆盖断言、集合/字符串、errx 结构化断言、JSON 语义相等、表格驱动、
  输出捕获、临时环境与并发辅助；
- 核心零第三方依赖（errx 断言子能力自 v0.2.0 起依赖 errx）；
- 消除家族各库重复编写 fake/helper 的问题。

非目标：

- 完整测试框架（不替代 testing）；
- 插件式比较器系统（pre-1.0 按需评估）；
- 测试报告生成。

## 2. 核心模型

```text
testx
├── TB 接口（Helper / Errorf / Fatalf 子集）
├── 核心断言：Equal / NotEqual / True / False / Nil / NotNil
│             Error / NoError / ErrorIs / Empty / NotEmpty / Len
├── errx 断言（v0.2.0）：ErrCode / ErrKind / ErrFields
├── JSON 语义相等（v0.2.0）：JSONEqual
├── 集合/字符串（v0.3.0）：Contains / NotContains / Subset / ElementsMatch
├── 表格驱动与近似（v0.4.0）：RunCases / Approx / Panics / NotPanics
└── 辅助（v0.5.0）：CaptureStdout / CaptureStderr / TempEnv / Concurrently
```

## 3. 失败语义

- 断言失败调用 `t.Errorf`，测试继续执行，便于收集同一用例的全部失败；
- 误用（如 `Len` 作用于不支持的类型）调用 `t.Fatalf`，立即终止；
- 所有断言先调用 `t.Helper()`，失败定位到业务测试代码行。

## 4. 消息约定

- 全部使用简体中文；
- 相等类消息输出"实际/期望"两行并附带类型；容器输出长度与元素；
- errx 断言输出错误码、分类与字段。

## 5. 版本与兼容

- 语义化版本；pre-1.0 按路线图演进；
- 每版完成即发布 tag，CI 全绿后 Release 自动生成；
- v1.0.0 是否发布由维护者决定，testx 只推进到 1.0 候选即停。
