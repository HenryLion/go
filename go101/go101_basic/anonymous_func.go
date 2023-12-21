package main

func main() {
	x, y := func() (int, int) {
		println("This function has no parameters.")
		return 3, 4
	}()
	println(x, y)

	func(a, b int) {
		println("a*a + b*b = ", a*a+b*b)
	}(x, y)

	func(x int) {
		println("x*x + y*y =", x*x+y*y)
	}(y)

	func() {
		println("x*x + y*y = ", x*x+y*y)
	}()
}
