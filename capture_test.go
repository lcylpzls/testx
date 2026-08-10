package testx

import (
	"fmt"
	"os"
	"sync"
	"testing"
)

func TestCaptureStdout(t *testing.T) {
	got := CaptureStdout(func() {
		fmt.Print("标准输出")
	})
	if got != "标准输出" {
		t.Fatalf("捕获输出不匹配：%q", got)
	}
	if os.Stdout == nil {
		t.Fatal("捕获后应恢复 os.Stdout")
	}
}

func TestCaptureStderr(t *testing.T) {
	got := CaptureStderr(func() {
		fmt.Fprint(os.Stderr, "错误输出")
	})
	if got != "错误输出" {
		t.Fatalf("捕获输出不匹配：%q", got)
	}
}

func TestCaptureStdoutPanic(t *testing.T) {
	old := os.Stdout
	var got any
	func() {
		defer func() { got = recover() }()
		CaptureStdout(func() {
			panic("崩溃")
		})
	}()
	if got != "崩溃" {
		t.Fatalf("panic 应向外传播：%v", got)
	}
	if os.Stdout != old {
		t.Fatal("panic 后应恢复 os.Stdout")
	}
}

func TestCaptureStdoutLargeOutput(t *testing.T) {
	// 超出管道缓冲区（64 KiB）的输出不应阻塞。
	data := make([]byte, 256*1024)
	for i := range data {
		data[i] = 'a'
	}
	got := CaptureStdout(func() {
		fmt.Print(string(data))
	})
	if len(got) != len(data) {
		t.Fatalf("大输出捕获不完整：%d != %d", len(got), len(data))
	}
}

func TestTempEnv(t *testing.T) {
	key := "TESTX_TEMP_ENV"
	if err := os.Setenv(key, "旧值"); err != nil {
		t.Fatalf("预置环境变量失败：%v", err)
	}
	t.Cleanup(func() {
		if os.Getenv(key) != "旧值" {
			t.Errorf("清理后环境变量未恢复：%q", os.Getenv(key))
		}
	})
	TempEnv(t, key, "新值")
	if os.Getenv(key) != "新值" {
		t.Fatalf("环境变量未生效：%q", os.Getenv(key))
	}
}

func TestTempEnvUnset(t *testing.T) {
	key := "TESTX_TEMP_UNSET"
	_ = os.Unsetenv(key)
	t.Cleanup(func() {
		if _, ok := os.LookupEnv(key); ok {
			t.Errorf("清理后环境变量应未设置")
		}
	})
	TempEnv(t, key, "值")
	if os.Getenv(key) != "值" {
		t.Fatalf("环境变量未生效：%q", os.Getenv(key))
	}
}

func TestTempEnvMisuse(t *testing.T) {
	tb := &fakeTB{}
	TempEnv(tb, "KEY")
	if len(tb.fatals) != 1 {
		t.Fatalf("奇数参数应 Fatalf：%v", tb.fatals)
	}
	tb2 := &fakeTB{}
	TempEnv(tb2, "A=B", "x")
	if len(tb2.fatals) != 1 {
		t.Fatalf("非法键应 Fatalf：%v", tb2.fatals)
	}
}

func TestConcurrently(t *testing.T) {
	var mu sync.Mutex
	count := 0
	Concurrently(t, 8, func() {
		mu.Lock()
		count++
		mu.Unlock()
	})
	if count != 8 {
		t.Fatalf("期望执行 8 次，实际 %d", count)
	}
	tb := &fakeTB{}
	Concurrently(tb, 0, func() {})
	if len(tb.fatals) != 1 {
		t.Fatalf("并发数 < 1 应 Fatalf：%v", tb.fatals)
	}
}
