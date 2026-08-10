package core

// TB 是断言所需的最小测试接口。
// *testing.T 天然满足；库自身测试可用替身验证失败路径。
type TB interface {
	// Helper 标记辅助函数，失败定位到业务测试代码行。
	Helper()
	// Errorf 记录失败但不终止测试。
	Errorf(format string, args ...any)
	// Fatalf 记录失败并立即终止（仅用于断言误用）。
	Fatalf(format string, args ...any)
	// Cleanup 注册测试结束时的清理函数。
	Cleanup(func())
}
