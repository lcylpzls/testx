// Package logx 提供 logx.Logger 的测试捕获与断言辅助。
package logx

import (
	"context"
	"fmt"
	"sync"

	core "github.com/lcylpzls/logx"
)

// Level 是 core.Level 的便捷别名。
type Level = core.Level

// 日志级别常量（与 core 对齐）。
const (
	OffLevel   = core.OffLevel
	DebugLevel = core.DebugLevel
	InfoLevel  = core.InfoLevel
	WarnLevel  = core.WarnLevel
	ErrorLevel = core.ErrorLevel
	PanicLevel = core.PanicLevel
	FatalLevel = core.FatalLevel
)

// Entry 是捕获到的一条日志。
type Entry struct {
	// Level 日志级别。
	Level core.Level
	// Message 日志消息（格式化变体已展开）。
	Message string
	// Fields 结构化字段。
	Fields core.FieldGroup
}

// Logger 是捕获日志的 logx.Logger 实现，可注入被测代码。
type Logger struct {
	mu      sync.Mutex
	entries []Entry
	base    core.FieldGroup
}

// New 创建日志捕获器。
func New() *Logger {
	return &Logger{}
}

// Snapshot 返回已捕获日志的副本。
func (l *Logger) Snapshot() []Entry {
	l.mu.Lock()
	defer l.mu.Unlock()
	out := make([]Entry, len(l.entries))
	copy(out, l.entries)
	return out
}

// IsDebugEnabled 恒返回 true。
func (l *Logger) IsDebugEnabled() bool {
	return true
}

func (l *Logger) Debug(msg string, fields core.FieldGroup) { l.record(core.DebugLevel, msg, fields) }
func (l *Logger) Info(msg string, fields core.FieldGroup)  { l.record(core.InfoLevel, msg, fields) }
func (l *Logger) Warn(msg string, fields core.FieldGroup)  { l.record(core.WarnLevel, msg, fields) }
func (l *Logger) Error(msg string, fields core.FieldGroup) { l.record(core.ErrorLevel, msg, fields) }

func (l *Logger) Panic(msg string, fields core.FieldGroup) {
	l.record(core.PanicLevel, msg, fields)
	panic(msg)
}

func (l *Logger) Fatal(msg string, fields core.FieldGroup) {
	l.record(core.FatalLevel, msg, fields)
	l.SafeExit(nil)
}

func (l *Logger) Debugf(format string, args ...any) {
	l.record(core.DebugLevel, fmt.Sprintf(format, args...), core.FieldGroup{})
}
func (l *Logger) Infof(format string, args ...any) {
	l.record(core.InfoLevel, fmt.Sprintf(format, args...), core.FieldGroup{})
}
func (l *Logger) Warnf(format string, args ...any) {
	l.record(core.WarnLevel, fmt.Sprintf(format, args...), core.FieldGroup{})
}
func (l *Logger) Errorf(format string, args ...any) {
	l.record(core.ErrorLevel, fmt.Sprintf(format, args...), core.FieldGroup{})
}

func (l *Logger) Panicf(format string, args ...any) {
	msg := fmt.Sprintf(format, args...)
	l.record(core.PanicLevel, msg, core.FieldGroup{})
	panic(msg)
}

func (l *Logger) Fatalf(format string, args ...any) {
	l.record(core.FatalLevel, fmt.Sprintf(format, args...), core.FieldGroup{})
	l.SafeExit(nil)
}

// WithContext 返回携带上下文的副本（捕获语义不变）。
func (l *Logger) WithContext(ctx context.Context) core.Logger { return l }

// WithField 返回携带默认字段的副本。
func (l *Logger) WithField(key string, val any) core.Logger {
	return &Logger{
		base: mergeFields(l.base, core.Fields(core.Any(key, val))),
	}
}

// Sync 无操作。
func (l *Logger) Sync() error { return nil }

// Close 无操作。
func (l *Logger) Close() error { return nil }

// SafeExit 执行退出回调（测试中用于验证 Fatal 路径）。
func (l *Logger) SafeExit(exitFunc func()) {
	if exitFunc != nil {
		exitFunc()
	}
}

func (l *Logger) record(level core.Level, msg string, fields core.FieldGroup) {
	l.mu.Lock()
	l.entries = append(l.entries, Entry{
		Level:   level,
		Message: msg,
		Fields:  mergeFields(l.base, fields),
	})
	l.mu.Unlock()
}

// mergeFields 合并两组字段。
func mergeFields(a, b core.FieldGroup) core.FieldGroup {
	if a.Len() == 0 {
		return b
	}
	if b.Len() == 0 {
		return a
	}
	var fs []core.Field
	for _, g := range []core.FieldGroup{a, b} {
		for i := 0; i < g.Len(); i++ {
			fs = append(fs, g.At(i))
		}
	}
	return core.Fields(fs...)
}
