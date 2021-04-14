package main

import "fmt"

func main() {
	var sum = 0
	for i := 1; i <= 100; i++ {
		sum += i
	}
	fmt.Println(sum)

	var arr = [9]int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	for i := 0; i < len(arr); i++ {
		fmt.Print(arr[i])
	}
}
