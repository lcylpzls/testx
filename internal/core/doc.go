// Package testx 是轻量、零魔法的测试断言与辅助库。
//
// 核心断言基于可替换的 TB 接口（Helper/Errorf/Fatalf 子集），
// *testing.T 天然满足；失败消息统一使用简体中文。
//
// 典型用法：
//
//	func TestGreeting(t *testing.T) {
//	    testx.Equal(t, greeting("小明"), "你好，小明！")
//	}
package core
