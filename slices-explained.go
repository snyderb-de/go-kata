package main

import "fmt"

func main() {

	// var ages [3]int = [3]int{20, 25, 30}
	var ages = [3]int{20, 25, 30} // type inference

	var names = [4]string{"Yoshi", "Mario", "Peach", "Bowser"}
	names[1] = "Luigi"

	fmt.Println(ages, len(ages))
	fmt.Println(names, len(names))

	// Slices
	var scores = []int{100, 50, 60} // no length means slice
	scores[2] = 25
	scores = append(scores, 85) // append returns a new slice

	fmt.Println(scores, len(scores))
}
