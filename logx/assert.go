package logx

import (
	"strings"

	"github.com/lcylpzls/testx"
)

// AssertLogged 断言捕获日志中存在指定级别与消息的条目。
func AssertLogged(t testx.TB, l *Logger, level Level, message string) {
	t.Helper()
	for _, e := range l.Snapshot() {
		if e.Level == level && e.Message == message {
			return
		}
	}
	t.Errorf("期望存在日志 %s %q，实际 %v", level, message, l.Snapshot())
}

// AssertLoggedContains 断言捕获日志中存在指定级别且消息包含子串的条目。
func AssertLoggedContains(t testx.TB, l *Logger, level Level, substring string) {
	t.Helper()
	for _, e := range l.Snapshot() {
		if e.Level == level && strings.Contains(e.Message, substring) {
			return
		}
	}
	t.Errorf("期望存在日志 %s 包含 %q，实际 %v", level, substring, l.Snapshot())
}
