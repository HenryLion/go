package main

type T int

// 1, the arguments of a deferred function call are evaluated when
// the deferred call is pushed into the deferred call queue

// 2, method receiver arguments like normal arguments

func (t T) M(n int) T {
	print(n)
	return t
}

func main() {
	var t T
	// when M(2) pushed into the deferred queue, t.M(1) has been evaluated
	defer t.M(1).M(2)
	t.M(3).M(4)
	// the result is 1342
}
