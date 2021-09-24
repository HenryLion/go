package main
import (
	"fmt"
)

type Point struct {
	x,y int
}

type PointNew struct {
	X, Y int
}

type Circle struct {
	Center PointNew
	Radius int
}

type Whell struct {
	circle Circle
	Spokes int
}

func main () {
	// create a struct
	p1 := Point{1,2}
	fmt.Println(p1) //{1 2}

	// create a struct with fields names
	p2 := Point{y:3, x:7}
	fmt.Println(p2) //{7 3}

	// create a pointer to struct
	pp1 := &Point{1,3}
	fmt.Println(pp1) // &{1 3}

	// create a pointer to struct
	pp := new(Point)
	*pp = Point{1,2}
	fmt.Println(pp) //&{1 2}
	fmt.Println(*pp) // {1 2}

	// pp = new(Point{2,4})  // complier error: Point{...} is not a type

	// create a wheel
	var w Whell
	w.circle.Center.X = 8
	w.circle.Center.Y = 8
	w.circle.Radius = 10
	w.Spokes = 20
	fmt.Println(w)

	// w1 := Whell {8,8,10,20} // compile error
	// w2 := Whell {X: 8, Y:8, Radius: 5, Spokes: 20} // compile error
	w3 := Whell {Circle{PointNew{8,8}, 5}, 20}
	fmt.Println(w3)
	fmt.Printf("%#v\n", w3)
}