package main

import "fmt"

func main() {
	//s := 10
	//ch := 'a'
	//b := true
	//fmt.Printf("%d, %c\n", s, ch)
	//fmt.Println(b, !b)
	//var str string = "abcdefg"
	//var str2 = "adsdasd"
	//fmt.Println(str)
	//fmt.Println(str)
	//fmt.Println(&str)
	//fmt.Println(&str2)
	//str += "  " + str2
	//fmt.Println(str, str[0], str[len(str)-1])
	//for i := 0; i < len(str); i++ {
	//	fmt.Printf("%c", str[i])
	//}
	//fmt.Printf("\n")
	//var isused bool
	//fmt.Printf("%v\n", isused)
	//s := 10
	//ps := &s
	//fmt.Println(*ps)
	//s := 10
	//if s < 10 {
	//	fmt.Println("m")
	//} else if s < 20 {
	//	fmt.Println("mm")
	//} else {
	//	fmt.Println("mmm")
	//}
	//s := 10
	//a := 1
	//fmt.Println(&s)
	//ptr := &s
	//*ptr = 122
	//fmt.Println(&s, " ", ptr)
	//if a != 0 {
	//	fmt.Println(a)
	//}

	a := 10
	b := 20
	fmt.Println(a, b)
	a ^= b
	b ^= a
	a ^= b
	fmt.Println(a, b)

}
