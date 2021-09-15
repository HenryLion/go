package main
import (
	"fmt"
	"math"
	"math/cmplx"
)

func main () {
	fmt.Println(math.MaxFloat64)
	var f float32 = 16777217
	fmt.Println(f == f+1)
	const f1 = 6.269e-28
	fmt.Println(f1)
	fmt.Printf("%g - %g\n", f, f1)
	fmt.Printf("%v - %v\n", f, f1)
	fmt.Printf("%6.2f - %f\n", f, f1)

	for x := 0; x < 8; x++ {
		fmt.Printf ("x = %d e^ = %8.4f\n", x, math.Exp(float64(x)))
	}

	// complex numbers
	var x complex128 = complex(1,2)
	var y complex128 = complex(3,4)
	fmt.Println(x + y)
	fmt.Println(x * y)
	fmt.Println(real(x*y))
	fmt.Println(imag(x+y))

	// 'i' in complex
	a := 1i
	b := 1i
	fmt.Println(a*b)

	// cmplx package
	fmt.Println(cmplx.Sqrt(-1))

	//bool

	var t bool = true
	// var t1 = int(t)  /* can not do this. error */
	fmt.Println(t)
}