package core

import (
	"io"
	"os"
	"sync"
)

// CaptureStdout 执行 fn 并捕获写入 os.Stdout 的内容。
// 仅捕获使用 os.Stdout 全局变量的输出，不支持子进程；不可用于并行测试。
func CaptureStdout(fn func()) string {
	return captureOutput(&os.Stdout, fn)
}

// CaptureStderr 执行 fn 并捕获写入 os.Stderr 的内容；限制同 CaptureStdout。
func CaptureStderr(fn func()) string {
	return captureOutput(&os.Stderr, fn)
}

// captureOutput 临时替换输出流，执行 fn 后恢复并返回捕获内容。
// os.Pipe 创建失败时视为空输出（该失败在常规环境不可达）。
func captureOutput(stream **os.File, fn func()) string {
	r, w, _ := os.Pipe()
	old := *stream
	*stream = w
	done := make(chan string, 1)
	go func() {
		data, _ := io.ReadAll(r)
		done <- string(data)
	}()
	defer func() {
		*stream = old
		w.Close()
		r.Close()
	}()
	fn()
	*stream = old
	w.Close()
	return <-done
}

// TempEnv 设置环境变量，并在测试结束时恢复原值（不存在的键恢复为未设置）。
// envs 必须为成对的 KEY/VALUE；键含 "=" 或 NUL 时按误用 Fatalf。
func TempEnv(t TB, envs ...string) {
	t.Helper()
	if len(envs)%2 != 0 {
		t.Fatalf("TempEnv 需要成对的 KEY/VALUE，得到 %d 个参数", len(envs))
		return
	}
	for i := 0; i < len(envs); i += 2 {
		key, val := envs[i], envs[i+1]
		prev, existed := os.LookupEnv(key)
		if err := os.Setenv(key, val); err != nil {
			t.Fatalf("TempEnv 设置环境变量 %q 失败：%v", key, err)
			return
		}
		t.Cleanup(func() {
			if existed {
				_ = os.Setenv(key, prev)
			} else {
				_ = os.Unsetenv(key)
			}
		})
	}
}

// Concurrently 并发运行 fn 共 n 次并等待全部完成；用于竞态检测。
// n 必须 >= 1；fn 必须并发安全。
func Concurrently(t TB, n int, fn func()) {
	t.Helper()
	if n < 1 {
		t.Fatalf("Concurrently 的并发数必须 >= 1，得到 %d", n)
		return
	}
	var wg sync.WaitGroup
	wg.Add(n)
	for i := 0; i < n; i++ {
		go func() {
			defer wg.Done()
			fn()
		}()
	}
	wg.Wait()
}
