package main

import (
	"fmt"
	"reflect"
)

type T []interface{ m() }

func (T) m() {}

type F func(string, int) bool

func (f F) m(s string) bool {
	return f(s, 32)
}
func (f F) M() {}

type I interface {
	m(s string) bool
	M()
}

func main() {
	type A = [16]int16
	var c <-chan map[A][]byte
	tc := reflect.TypeOf(c)
	fmt.Println(tc.Kind())
	fmt.Println(tc.ChanDir())

	tm := tc.Elem()
	fmt.Println(tm)
	ta, tb := tm.Key(), tm.Elem()
	fmt.Println(ta, tb)
	fmt.Println(tm.Kind(), ta.Kind(), tb.Kind())

	// use the Elem get the base type of pointer type

	fmt.Println("==============")
	tp := reflect.TypeOf(new(interface{}))
	tt := reflect.TypeOf(T{})
	fmt.Println(tp.Kind(), tt.Kind())

	ti := tp.Elem()
	tim := tt.Elem()
	fmt.Println(ti.Kind(), tim.Kind()) // interface interface
	fmt.Println(tt.Implements(tim))

	// struct and func typeof
	fmt.Println("==============")
	var x struct {
		F F
		i I
	}
	tx := reflect.TypeOf(x)
	fmt.Println(tx.Kind())        // struct
	fmt.Println(tx.NumField())    // 2
	fmt.Println(tx.Field(0).Name) // F
	fmt.Println(tx.Field(0).PkgPath)
	fmt.Println(tx.Field(1).PkgPath) // main

	tf, tii := tx.Field(0).Type, tx.Field(1).Type
	fmt.Println(tf.Kind())               // func
	fmt.Println(tf.IsVariadic())         // false
	fmt.Println(tf.NumIn(), tf.NumOut()) // 2 1

	t0, t1, t2 := tf.In(0), tf.In(1), tf.Out(0)
	fmt.Println(t0.Kind(), t1.Kind(), t2.Kind()) // string int bool

	fmt.Println(tf.NumMethod(), tii.NumMethod()) // 1 2 for non-interface type it returns the number of exported methods
	fmt.Println(tf.Method(0).Name)               // M
	fmt.Println(ti.Method(1).Name)               // m
	_, ok1 := tf.MethodByName("m")
	_, ok2 := tii.MethodByName("m")
	fmt.Println(ok1, ok2) // false true

}
