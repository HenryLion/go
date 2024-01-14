package main

import (
	"fmt"
	"reflect"
)

type T1 struct {
	X    int  `max:"99" min:"0", default:"0"`
	Y, Z bool `optional:"yes1"`
}

func main() {
	t := reflect.TypeOf(T1{})
	x := t.Field(0).Tag
	fmt.Println(x) // max:"99" min:"0", default:"0"
	y := t.Field(1).Tag
	z := t.Field(2).Tag
	fmt.Println(reflect.TypeOf(x)) // reflect.StructTag
	v, present := x.Lookup("max")
	fmt.Println(len(v), present)      // 2 true, 2 means len("99")
	fmt.Println(x.Get("max"))         // 99
	fmt.Println(x.Lookup("optional")) // false
	fmt.Println(y.Lookup("optional")) // yes1 true
	fmt.Println(z.Lookup("optional")) // yes1 true
}
