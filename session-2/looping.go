package main

import "fmt"

func main() {
	// iterator akan bekerja sampai limit
	for i := 1; i <= 100; i++ {
		fmt.Println("Angka", i)
	}
}