package main

func main() {
	var target = 64
	if target >= 80 {
		print("A")
	} else if target <= 80 && target >= 60 {
		print("B")
	} else {
		print("C")
	}
}
