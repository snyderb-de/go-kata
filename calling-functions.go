package main

import "fmt"

func learning(userName string) string {
	return userName + " is learning how to call functions!"
}

func main() {
	// Do not change this two lines of code
	var userName string
	fmt.Scanf("%s", &userName)

	fmt.Println(learning(userName))
}
