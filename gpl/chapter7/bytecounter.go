package main
import "fmt"
import "sort"

type ByteCounter int

func (c *ByteCounter) Write (p []byte) (int, error) {
	*c += ByteCounter(len(p))
	return len(p), nil
}

type NameWriter string

func (s *NameWriter) Write (p []byte) (int, error) {
	*s = NameWriter(string(p))
	return len(p), nil
}


type Add interface {
	add (a int, b int) int
}

type Sub interface {
	sub (a int, b int) int
}

type AddSub interface {
	Add
	Sub
}

func (c *ByteCounter) add (a int, b int) int {
	return a + b
}

func (c *ByteCounter) sub (a int, b int) int {
	return a - b
}


func print_any (any interface{}) {
	fmt.Println(any)
}

func main () {

	// a type can pass to a interface when the type has all the method the interface has
	var face1 AddSub
	var c0 *ByteCounter
	face1 = c0
	fmt.Println(face1.add(3,4))
	fmt.Println(face1.sub(3,4))


	var c ByteCounter
	c.Write([]byte("hellio"))
	fmt.Println(c) // method 的用法

	c = 0
	var name = "pencil"
	fmt.Fprintf(&c, "hello, %s", name)
	fmt.Println(c)


	var s NameWriter
	s.Write([]byte("I want you know."))
	fmt.Println(s)

	s = ""
	var sentence = "today is a good day."
	fmt.Fprintf(&s, "Meimei say: %s", sentence)
	fmt.Println(s)

	// a empty interface can receive any type
	var any interface {}
	any = true
	fmt.Println(any)
	any = 12.34
	fmt.Println(any)
	any = "hello"
	fmt.Println(any)
	any = map[string]int{"one":1}
	fmt.Println(any)

	print_any(111)
	print_any("hello")
	print_any(any)

	// sort package
	names := [] string {"zhang","wang","li","zhao"}
	fmt.Println(names)

	sort.Strings(names)
	fmt.Println(names)

}