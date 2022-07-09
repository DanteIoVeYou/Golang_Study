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

//package main
//
//import (
//	"lib1"
//	"lib2"
//)
//
//func main() {
//	lib1.Lib1Test()
//	lib2.Lib2Test()
//}

// ==========================================================
// 给包起别名：别名 包名
// 导入了包不想使用，包起别名，加  下划线 空格   _ "name"
// . 报名： 不需要 包.方法名，直接方法名
// ==========================================================

// 指针
//func swap(a *int, b *int) {
//	tmp := *a;
//	*a = *b;
//	*b = tmp;
//}

// ==========================================================

// defer

//package main
//
//import "fmt"
//
//func deferFunc() int {
//	fmt.Println("defer call")
//	return 1
//}
//func returnFunc() int {
//	fmt.Println("return call")
//	return 1
//}
//func Testdefer() int {
//	defer deferFunc()
//	return returnFunc()
//}
//func main() {
//	//fmt.Println("1")
//	//fmt.Println("2")
//	//fmt.Println("3")
//	//defer fmt.Println("1")
//	//defer fmt.Println("2")
//	//defer fmt.Println("3")
//	Testdefer()
//}

// ==========================================================

// 数组

//package main
//
//import "fmt"
//
//func print(arr []int) {
//	// 下划线表示匿名变量
//	arr[0] = 100
//	for _, value := range arr {
//		fmt.Println("value = ", value)
//
//	}
//}
//func main() {
//	var arr = [10]int{1,2,3,4}
//	for i := 0; i < 10; i++ {
//		fmt.Println(arr[i])
//	}
//	for index,value := range arr {
//		fmt.Println("index: ", index, "value: ", value)
//	}
//	// 查看数组的数据类型
//	fmt.Printf("type of arr is %T\n", arr)
//	// 动态数组，切片
//	// 传参是引用传递，而不是值传递
//	arr1 := []int{1,2,3,4,5}
//	fmt.Printf("type of arr1 is %T\n", arr1)
//	print(arr1)
//}

// ==========================================================

// slice的声明
//package main
//
//import "fmt"
//
//func main() {
//	// 1. 声明slice1是一个切片，并且初始化，且长度是3
//	slice1 := []int{1,2,3}
//	fmt.Printf("len = %d, value: %v\n", len(slice1), slice1)
//	// 2.用make对slice分配空间
//	var slice2 []int
//	slice2 = make([]int, 3)
//	// 3.声明时用make
//	var slice3 []int = make([]int, 3)
//	// 4. := + make
//	slice4 := make([]int, 3)
//	fmt.Printf("len = %d, valye: %v\n", len(slice2), slice2)
//	fmt.Printf("len = %d, valye: %v\n", len(slice3), slice3)
//	fmt.Printf("len = %d, valye: %v\n", len(slice4), slice4)
//	// 判断slice是否为空
//	var slice5 []int
//	if slice5 == nil {
//		fmt.Println("slice5 is nil")
//	} else {
//		fmt.Println("slice5 is not nil")
//	}
//}

// ==========================================================

// slice的追加与截取
//package main
//
//import "fmt"
//
//func main() {
//	slice1 := make([]int, 3, 5)
//	fmt.Printf("slice1 len = %d, cap = %d, value: %v\n", len(slice1), cap(slice1), slice1)
//	slice1 = append(slice1, 1)
//	fmt.Printf("slice1 len = %d, cap = %d, value: %v\n", len(slice1), cap(slice1), slice1)
//	slice1 = append(slice1, 1)
//	fmt.Printf("slice1 len = %d, cap = %d, value: %v\n", len(slice1), cap(slice1), slice1)
//	// 截取的slice2和slice1指向同一块空间
//	slice2 := slice1[:3]
//	slice2[0] = 3
//	fmt.Printf("slice2 len = %d, cap = %d, value: %v\n", len(slice2), cap(slice2), slice2)
//	fmt.Printf("slice1 len = %d, cap = %d, value: %v\n", len(slice1), cap(slice1), slice1)
//	// copy函数提供深拷贝
//	slice3 := make([]int, 5)
//	copy(slice3, slice2)
//	slice3[0] = 999
//	fmt.Printf("slice3 len = %d, cap = %d, value: %v\n", len(slice3), cap(slice3), slice3)
//	fmt.Printf("slice2 len = %d, cap = %d, value: %v\n", len(slice2), cap(slice2), slice2)
//
//}

// ==========================================================

// map

//package main
//
//import "fmt"
//
//func main() {
//	// 声明map类型
//	var myMap map[string]string
//	if myMap == nil {
//		fmt.Println("nil")
//	}
//	// 使用map前需要先用make分配空间
//	myMap = make(map[string]string, 10)
//	myMap["one"] = "c++"
//	myMap["two"] = "java"
//	myMap["three"] = "python"
//	fmt.Println(myMap)
//
//	myMap2 := make(map[int]string)
//	myMap2[1] = "php"
//	fmt.Println(myMap2)
//
//	myMap3 := map[int]string {
//		1 : "go",
//		2 : "c#",
//	}
//	fmt.Println(myMap3)
//
//}

// ==========================================================

// map的使用方式
//package main
//
//import "fmt"
//func print(myMap map[string]string) {
//	// map是引用传递
//	for k,v := range myMap {
//		fmt.Println("key: ", k, "value: ", v)
//	}
//}
//
//func main() {
//	// 添加
//	myMap := make(map[string]string)
//	myMap["apple"] = "苹果"
//	myMap["banana"] = "香蕉"
//	myMap["pear"] = "梨"
//	// 遍历
//	for k,v := range myMap {
//		fmt.Println("key: ", k, "value: ", v)
//	}
//	// 删除
//	delete(myMap, "apple")
//	fmt.Println("================================")
//	// 修改
//	myMap["banana"] = "香蕉香蕉"
//	for k,v := range myMap {
//		fmt.Println("key: ", k, "value: ", v)
//	}
//	fmt.Println("================================")
//	print(myMap)
//}

// ==========================================================
// struct

//package main
//
//import "fmt"
//
//type Book struct {
//	name string
//	price float32
//}
//func changePrice(book *Book, price float32) {
//	(*book).price = price
//}
//func main() {
//	var book1 Book
//	book1.name = "哈利波特"
//	book1.price = 20.11
//	fmt.Println(book1)
//	changePrice(&book1, 30.22)
//	fmt.Println(book1)
//}

// ==========================================================
// 封装
// 如果类名首字母表示其他包也能访问
// 类属性首字母大写表示公有，小写表示公有

//package main
//
//import "fmt"
//func (this *Hero) SetName(name string) {
//	this.Name = name
//}
//func (this *Hero) GetName() string {
//	return this.Name
//}
//func (this *Hero) ShowName() {
//	fmt.Println(this.Name)
//}
//
//type Hero struct {
//	Name string
//	Ad int
//	Level int
//}
//func main() {
//	hero := Hero{"laa", 11,32}
//	hero.ShowName()
//	hero.SetName("zhangsan")
//	name := hero.GetName()
//	fmt.Println(name)
//	hero.ShowName()
//}
// ==========================================================
// 继承

//package main
//
//import "fmt"
//
//
//type Human struct {
//	Name string
//	Age int
//}
//type Superman struct {
//	Human
//	Level int
//}
//func (this *Human) Eat() {
//	fmt.Println("human eat...")
//}
//func (this *Human) Play() {
//	fmt.Println("human play...")
//}
//func (this *Superman) Eat() {
//	fmt.Println("superman eat...")
//}
//func (this *Superman) Play() {
//	fmt.Println("superman play...")
//}
//func (this *Superman) Fly() {
//	fmt.Println("supeman fly...")
//}
//func main() {
//	var human1 Human = Human{"a", 12}
//	human1.Eat()
//	human1.Play()
//	var superman1 Superman = Superman{Human{"b", 12}, 100}
//	superman1.Eat()
//	superman1.Play()
//	superman1.Fly()
//}

// ==========================================================
// 多态

//package main
//
//import "fmt"
//
//type AnimalIF interface {
//	Sleep()
//	GetColor() string
//	GetType() string
//}
//type Cat struct {
//	color string
//}
//func (this *Cat) Sleep() {
//	fmt.Println("Cat sleep...")
//}
//func (this *Cat) GetColor() string {
//	return this.color
//}
//func (this *Cat) GetType() string {
//	return "Cat"
//}
//type Dog struct {
//	color string
//}
//func (this *Dog) Sleep() {
//	fmt.Println("Dog sleep...")
//}
//func (this* Dog) GetColor() string {
//	return this.color
//}
//func (this *Dog) GetType() string {
//	return "Dog"
//}
//func PrintAnimal(animal AnimalIF) {
//	animal.Sleep()
//	fmt.Println("color: ", animal.GetColor())
//	fmt.Println("type: ", animal.GetType())
//}
//func main() {
//
//	PrintAnimal(&Cat{"Blue"})
//	fmt.Println("====================================")
//	PrintAnimal(&Dog{"Yellow"})
//
//}

// ==========================================================
// interface{} 空接口 和 类型断言
// interface是万能数据类型

//package main
//
//import "fmt"
//
//func INTER(arg interface{}) {
//	fmt.Println("INTER called...")
//	fmt.Println(arg)
//	// 断言，判断数据类型
//	value, ok := arg.(string)
//	if ok == true {
//		fmt.Println("string", value)
//	} else {
//		fmt.Println("not string")
//	}
//	fmt.Println("===================================")
//}
//func main() {
//	INTER(12)
//	INTER("hello")
//}

// ==========================================================
//
// 变量 = 类型 + 值

//package main
//
//func main() {
//
//}
// ==========================================================
// reflect
//package main
//
//import (
//	"fmt"
//	"reflect"
//)
//
//type User struct {
//	Id int
//	Name string
//	Age int
//}
//
//func RefNum(arg interface{}) {
//	inputType := reflect.TypeOf(arg)
//	inputValue := reflect.ValueOf(arg)
//	fmt.Println("type is ", inputType)
//	fmt.Println("value is ", inputValue)
//	for i := 0; i < inputType.NumField(); i++ {
//		field := inputType.Field(i)
//		value := inputValue.Field(i).Interface()
//		fmt.Printf("%s: %v = %v\n", field.Name, field.Type, value)
//	}
//}
//
//func main() {
//	user := User{1, "12", 133}
//	RefNum(user)
//}
// ==========================================================
// Tag

//package main
//
//type resume struct {
//	Name string `info:"name" doc:"我的名字"`
//	Sex string `info:"sex"`
//}
//func main() {
//
//}
// ==========================================================
//goroutine

//package main
//
//import (
//	"fmt"
//	"time"
//)
//
//func newTask() {
//	i := 0
//	for {
//		i++
//		fmt.Println("new task: ", i)
//		time.Sleep(1 * time.Second)
//	}
//}
//func main() {
//	// 创建一个goroutine执行newTask()方法
//	go newTask()
//	i := 0
//	for {
//		i++
//		fmt.Println("main goroutine i = ", i)
//		time.Sleep(1 * time.Second)
//	}
//
//}

package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	// 用go创建一个形参为空，返回值为空的一个函数
	go func() {
		defer fmt.Println("A.defer")
		func() {
			defer fmt.Println("B.defer")
			runtime.Goexit()
			fmt.Println("B")
		}()
		fmt.Println("A")
	}()
	for {
		time.Sleep(1 * time.Second)
	}
}



// ==========================================================
// ==========================================================
// ==========================================================
// ==========================================================
// ==========================================================
// ==========================================================
