package main

import "fmt"

type Element struct {
	Name         string
	Symbol       string
	AtomicNumber int
}

func main() {
	// below is the original 'sodium' struct
	sodium := Element{
		Name:         "Sodium",
		Symbol:       "Na",
		AtomicNumber: 11,
	}

	// below are the various copies of 'sodium' Professor Ross created!
	sodium1 := Element{Name: "Sodium", Symbol: "Na"}
	sodium2 := Element{Name: "Sodium", Symbol: "Na", AtomicNumber: 11}
	sodium3 := Element{Symbol: "Na", AtomicNumber: 11}

	// compare the structs and print the results
	fmt.Println(sodium == sodium1)
	fmt.Println(sodium == sodium2)
	fmt.Println(sodium == sodium3)
	// compare sodium1 and sodium2 and 3 to each other and print the results
	fmt.Println(sodium1 == sodium2)
	fmt.Println(sodium1 == sodium3)
	fmt.Println(sodium2 == sodium3)

}
