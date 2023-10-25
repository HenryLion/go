package main
import "fmt"
import "time"
import "log"

func defer_func(msg string) {
	fmt.Println(msg)
}

func defer_func2(msg string) {
	fmt.Println(msg)
}

func trace (msg string) func() {
	start := time.Now()
	log.Printf("entry %s", msg)
	return func() {
		log.Printf("exit %s (%s)", msg, time.Since(start))
	}
}

func bigSlowOperation() {
	defer trace("bigSlowOperator")() // 其实真正defer的是trace返回的那个函数
	time.Sleep(3 * time.Second)
}

// // this is func defination
// func() { log.Printf("%s", "hello,world") }

// // this is func call
// func() { log.Printf("%s", "hello,world") } ()

func double(x int) (result int) {
	defer func() {
		fmt.Printf("double(%d) = %d\n", x, result)
	} ()
	return x + x
}

func main() {
	// defer 的基本特性
	defer defer_func("马上结束")
	defer defer_func("马上结束niao")
	fmt.Println("main begin.")

	// ..
	bigSlowOperation()

	// defer执行的时候其函数的result值已经更新了
	double(4)


}