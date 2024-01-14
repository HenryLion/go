package main

import (
	"fmt"
	"reflect"
)

func main() {
	n := 123
	p := &n
	vp := reflect.ValueOf(p)
	fmt.Println(p, vp)
	fmt.Println(vp.CanSet(), vp.CanAddr()) // false false
	vn := vp.Elem()
	fmt.Println(vn.CanSet(), vn.CanAddr()) // true true
	vn.Set(reflect.ValueOf(789))
	fmt.Println(n) // 789

	// non-exported fields of struct values can't be modified by reflections
	var s struct {
		X interface{}
		y interface{}
	}
	vp1 := reflect.ValueOf(&s)
	vs := reflect.Indirect(vp1) // equal to vs := vp.Elem()
	vx, vy := vs.Field(0), vs.Field(1)
	fmt.Println(vx.CanSet(), vx.CanAddr()) // true  true
	fmt.Println(vy.CanSet(), vy.CanAddr()) // false true
	vb := reflect.ValueOf(123)
	vx.Set(vb)
	fmt.Println(s)                      // {123 <nil>}
	fmt.Println(vx.IsNil(), vy.IsNil()) // false true

	var z = 123
	var y = &z
	var x interface{} = y
	v := reflect.ValueOf(&x)
	vx1 := v.Elem()
	vy1 := vx1.Elem()
	vz1 := vy1.Elem()
	vz1.Set(reflect.ValueOf(789))
	fmt.Println(z) //789
}
