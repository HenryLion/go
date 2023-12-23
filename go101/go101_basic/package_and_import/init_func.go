package main

import (
	"fmt"
	"runtime"
)

//import (
//	"database/sql"
//	_ "github.com/go-sql-driver/mysql" 导入此匿名包是为了初始化mysql驱动，以便database/sql包可以调用sql.Open连接数据库
//)

func init() {
	fmt.Println("hi,", bob)
}

func main() {
	fmt.Println("bye")
	fmt.Println(runtime.NumCPU())
}

func init() {
	fmt.Println("hello,", smith)
}

func titledName(who string) string {
	return "Mr. " + who
}

var bob, smith = titledName("Bob"), titledName("Smith")
