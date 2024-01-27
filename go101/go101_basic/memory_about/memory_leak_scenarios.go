package main

// 内存泄漏的场景
// 1、 取一个base字符串的子串时，
// base串已经不使用，但是由于子串引用了其中一小块内存
// 使得base串不得释放

var s0 string // a package-level variable

// A demo purpose function.
func f11(s1 string) {
	s0 = s1[:50]
	// Now, s0 shares the same underlying memory block
	// with s1. Although s1 is not alive now, but s0
	// is still alive, so the memory block they share
	// couldn't be collected, though there are only 50
	// bytes used in the block and all other bytes in
	// the block become unavailable.
}

func createStringWithLengthOnHeap(len int) string {
	return "this is a very long string about 1 M contents"
}

func demo() {
	s := createStringWithLengthOnHeap(1 << 20) // 1M bytes
	f11(s)
}

// 2、 取slice的一部分也会存在内存泄漏，原因基本同上

// 3、forever blocking goroutine 或许会永久占用一些资源

// 4、没有正确stop的time.Ticker

func main() {

}
