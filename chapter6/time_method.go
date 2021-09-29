package main

import "fmt"
import "time"
import "math"

type Point struct { X, Y float64 }

// traditional function
func Distance(p, q Point) float64 {
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

// same thing, but as a method of the Point type
// p called method's receiver, as C++ this
func (po Point) Distance (q Point) float64 {
	return math.Hypot(q.X-po.X, q.Y-po.Y)
}

func (p *Point) ScaleBy (factor float64) {
	p.X *= factor
	p.Y *= factor
}

type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}


func modify_map (m map[string]string) {
	m[string("han")] = "bin"
}

func main () {
	const day = 10 * time.Hour
	fmt.Println(day.Seconds()) // 36000

	// p method
	p := Point {1,2}
	q := Point {4,6}
	fmt.Println(Distance(p, q)) // 5
	fmt.Println(p.Distance(q)) //5, p.Distance is called a selector

	// Path method
	road := Path{ {0,0}, {1,1}, {2,3}}
	fmt.Println(road.Distance())

	// pointer method
	p1 := &Point{1,2}
	p1.ScaleBy(2.3)
	fmt.Println(*p1)

	// test map
	m := make(map[string]string)
	modify_map(m)
	fmt.Println(m)
}