package logx

import (
	"context"
	"errors"
	"testing"

	core "github.com/lcylpzls/logx"
)

func TestCaptureAndAssert(t *testing.T) {
	l := New()
	l.Info("服务启动", core.Fields(core.Int("port", 8080)))
	l.Debug("调试", core.FieldGroup{})
	l.Debugf("调试 %d", 1)
	l.Infof("信息 %d", 2)
	l.Warnf("警告 %d", 3)
	l.Errorf("错误 %d", 4)
	l.Warn("磁盘不足", core.FieldGroup{})
	l.Error("失败", core.Fields(core.Err(errors.New("boom"))))
	func() {
		defer func() { _ = recover() }()
		l.Panicf("恐慌 %d", 5)
	}()
	l.Fatalf("致命 %d", 6)

	AssertLogged(t, l, InfoLevel, "服务启动")
	AssertLoggedContains(t, l, DebugLevel, "调试")
	AssertLoggedContains(t, l, ErrorLevel, "失败")
	AssertLoggedContains(t, l, WarnLevel, "警告")
	AssertLoggedContains(t, l, PanicLevel, "恐慌")
	AssertLoggedContains(t, l, FatalLevel, "致命")
}

func TestAssertFailures(t *testing.T) {
	l := New()
	l.Info("其他消息", core.FieldGroup{})
	tb := &fakeTB{}
	AssertLogged(tb, l, InfoLevel, "不存在的消息")
	if !tb.failed() {
		t.Fatal("断言缺失应失败")
	}
	tb2 := &fakeTB{}
	AssertLoggedContains(tb2, l, WarnLevel, "不存在")
	if !tb2.failed() {
		t.Fatal("子串缺失应失败")
	}
}

func TestPanicAndFatal(t *testing.T) {
	l := New()
	func() {
		defer func() { _ = recover() }()
		l.Panic("恐慌", core.FieldGroup{})
	}()
	l.Fatal("致命", core.FieldGroup{})
	AssertLogged(t, l, PanicLevel, "恐慌")
	AssertLogged(t, l, FatalLevel, "致命")

	exited := false
	l.SafeExit(func() { exited = true })
	if !exited {
		t.Fatal("SafeExit 应执行回调")
	}
}

func TestWithFieldAndContext(t *testing.T) {
	l := New()
	if len(l.Snapshot()) != 0 {
		t.Fatal("新捕获器应为空")
	}
	derived := l.WithField("app", "demo")
	derived.Info("带字段", core.FieldGroup{})
	derived.Error("合并字段", core.Fields(core.String("env", "test")))
	_ = l.WithContext(context.Background())
	_ = l.IsDebugEnabled()
	_ = l.Sync()
	_ = l.Close()
	snap := derived.(*Logger).Snapshot()
	if len(snap) != 2 || snap[0].Message != "带字段" {
		t.Fatalf("派生日志不匹配：%+v", snap)
	}
	var hasField bool
	for i := 0; i < snap[0].Fields.Len(); i++ {
		if snap[0].Fields.At(i).Key == "app" {
			hasField = true
		}
	}
	if !hasField {
		t.Fatal("WithField 默认字段缺失")
	}
}

type fakeTB struct {
	errors []string
}

func (f *fakeTB) Helper()                           {}
func (f *fakeTB) Errorf(format string, args ...any) { f.errors = append(f.errors, format) }
func (f *fakeTB) Fatalf(format string, args ...any) {}
func (f *fakeTB) Cleanup(func())                    {}
func (f *fakeTB) failed() bool                      { return len(f.errors) > 0 }
