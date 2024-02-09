package main

import "fmt"

func main() {
	var firstName string = "my"
	const lastName string = "skill"

	firstName = "aku"
	lastName= "kemapuan"
	fmt.Print("halo ", firstName, lastName, "!\n")
}