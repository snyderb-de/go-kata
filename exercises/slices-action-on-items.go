package main

import "fmt"

func main() {
	// create vars and then scan them in and then create a slice from them
	var num1, num2, num3 int
	fmt.Scanln(&num1, &num2, &num3)
	intSlice := []int{num1, num2, num3}

	// Print each item in the slice multiplied by 2, on a new line
	for _, item := range intSlice {
		fmt.Println(item * 2)
	}

}
