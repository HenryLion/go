package main

import "fmt"

type Writer interface {
	write(buf []byte) (int, error)
}

type DummyWriter struct{}

func (DummyWriter) write(buf []byte) (int, error) {
	return len(buf), nil
}

func main() {

	//type assertion
	var x1 any = 123
	n, ok1 := x1.(int)
	fmt.Println(n, ok1)
	n = x1.(int)
	fmt.Println(n)

	a, ok1 := x1.(float64)
	fmt.Println(a, ok1)
	// a = x.(float64) // pay attention, this will panic

	var x interface{} = DummyWriter{}
	//var y any = "abc"
	var w Writer
	var ok bool

	w, ok = x.(Writer)
	fmt.Println(w, ok)
	x, ok = w.(interface{})
	fmt.Println(x, ok)
}
