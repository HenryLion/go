package main
import (
	"fmt"
)

const (
	e = 2.71828182845904523536028747135266249775724709369995957496696763
	pi = 3.14159265358979323846264338327950288419716939937510582097494459
)

const (
	sunday = 1
	monday
	tuesday
)


// const with iota like "enum"
const (
	one = iota
	another
	other
)

type Flags uint

const (
	Up Flags = 1 << iota
	Down
	Right
	Left
)

const (
	KB, MB, GB = 1e+03, 1e+06, 1e+09
)

func main () {
	fmt.Println(e)cd

	fmt.Println(monday,tuesday)
	fmt.Println(another,other)

	fmt.Println(Up, Down, Right, Left)

	fmt.Println(GB)
}