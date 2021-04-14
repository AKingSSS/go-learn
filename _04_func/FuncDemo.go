package main

import "fmt"

func max(num1, num2 int) int {
	if num1 > num2 {
		return num1
	} else {
		return num2
	}
}

func main() {
	var a = 4
	var b = 10
	fmt.Println(max(a, b))
}
