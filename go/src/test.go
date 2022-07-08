// package main

// import "fmt"

// func main() {
// 	fmt.Println("helloworld")
// }

// ==========================================================
//package main
//
//import "fmt"
//
//func main() {
//	// 方法一：声明一个变量，默认值是0
//	var a int
//	fmt.Println(a)
//	fmt.Printf("typeof a is %T\n", a)
//	// 方法二：赋初值
//	var b int = 100
//	fmt.Println(b)
//	fmt.Printf("typeof b is %T\n", a)
//	// 方法三：省略类型
//	var c = 100
//	fmt.Println(c)
//	fmt.Printf("typeof c is %T\n", a)
//	// 方法四：省去var，只能用在函数体内
//	d := "dasd"
//	fmt.Println(d)
//	fmt.Printf("typeof d is %T\n", d)
//	var (
//		aa int = 200
//		str string = "dasd"
//	)
//	fmt.Println(aa, " " ,str)
// ==========================================================
//package main
//
//import "fmt"
//
//const (
//	BEIJING = 10 * iota
//	SHANGHAI
//	SHENZHEN
//)
//func main() {
//	fmt.Println("Beijing = ", BEIJING)
//	fmt.Println("Shanghai = ", SHANGHAI)
//	fmt.Println("Shenzhen = ", SHENZHEN)
//}

// ==========================================================
// 函数多返回值的3种写法

//package main
//
//import "fmt"
//
//func foo1(a int, s string) (int, int) {
//	fmt.Println(a)
//	fmt.Println(s)
//	c := 1
//	d := 2
//	return c, d
//}
//// ret1，ret2属于形参，默认值是0
//func foo2(a string) (ret1 int, ret2 int) {
//	fmt.Println(a)
//	ret1 = 100
//	ret2 = 200
//	return ret1,ret2
//}
//func foo3(a string) (ret1, ret2 int) {
//	fmt.Println(a)
//	ret1 = 100
//	ret2 = 200
//	return ret1,ret2
//}
//func main() {
//	c, d := foo1(100, "hello")
//	fmt.Println(c)
//	fmt.Println(d)
//	foo2("hello")
//	foo3("world")
//}
// ==========================================================
// import导包路径和init方法

// 顺序
// import pkg1
// const
// var
// init()
// main()

package main

import (
	"lib1"
	"lib2"
)

func main() {
	lib1.Lib1Test()
	lib2.Lib2Test()
}

// ==========================================================
