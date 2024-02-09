package main

import "fmt"

func main() {
	var firstName string = "my"
	var lastName string = "skill"

	var numberA int = 4
	
	// syntax & untuk memanggil alamat memori.
	var numberB *int = &numberA

	// akan mengubah value numberA 
	*numberB = 8

	// memanggil alamat memori numberB 
	fmt.Print("halo ", firstName, lastName, numberB, "!\n")

	// tambahkan * untuk memanggil nilai numberB
	fmt.Print("halo ", firstName, lastName, *numberB, "!\n*")

	fmt.Print("halo ", firstName, lastName, numberA, "!\n*")

}