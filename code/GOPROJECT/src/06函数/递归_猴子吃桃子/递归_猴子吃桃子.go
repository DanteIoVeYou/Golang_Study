package main

import "fmt"

func eat(round int) int {
	if round < 10 {
		return 2 * (eat(round+1) + 1)
	} else {
		return 0
	}
}
func main() {
	fmt.Println(eat(1))
}
