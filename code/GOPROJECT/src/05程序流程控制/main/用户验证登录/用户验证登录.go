package main

import (
	"fmt"
)

func main() {
	var username string
	var passwd string
	var count int
	for count = 3; count > 0; count-- {
		fmt.Println("请输入用户名:>")
		fmt.Scanln(&username)
		fmt.Println("请输入密码:>")
		fmt.Scanln(&passwd)
		if username == "张无忌" && passwd == "888" {
			fmt.Println("登陆成功")
			break
		} else {
			fmt.Printf("你还有 %v 次机会\n", count-1)
		}
	}
	if count == 0 {
		fmt.Println("登陆失败")
	}
}
