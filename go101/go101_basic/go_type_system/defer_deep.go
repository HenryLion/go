package main

import "os"

type T11 int

// 1, the arguments of a deferred function call are evaluated when
// the deferred call is pushed into the deferred call queue

// 2, method receiver arguments like normal arguments

func (t T11) M(n int) T11 {
	print(n)
	return t
}

// 以下这个函数主要是想看下匿名函数的用法，匿名函数使得defer在每次loop中被执行一次
// 防止有太多文件需要关闭
func writeManyFiles(files []os.File) error {
	for _, file := range files {
		if err := func() error {
			f, err := os.Open(file.Name())
			if err != nil {
				return err
			}
			// the close method will be called at
			// the end of the current loop
			defer f.Close()
			_, err = f.WriteString("some content")
			if err != nil {
				return err
			}
			return f.Sync()
		}(); err != nil {
			return err
		}
	}
	return nil
}

func main() {
	var t T11
	// when M(2) pushed into the deferred queue, t.M(1) has been evaluated
	defer t.M(1).M(2)
	t.M(3).M(4)
	// the result is 1342
}
