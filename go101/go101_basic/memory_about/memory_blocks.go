package main

// go:noinline

func f(i int) byte {
	var a [1 << 20]byte
	return a[i]
}

func main() {
	var x int
	println(&x)
	f(100)
	println(&x)
}
