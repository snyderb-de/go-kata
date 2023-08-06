package main

import "fmt"

func main() {
	s := []int{2, 3, 5, 7, 11} // len: 5 cap: 5
	s1 := s[1:3]
	fmt.Println(s1) // [3 5]
}
