package main

import "fmt"

func main() {
	var point = 9
	switch point {
		case 10: 
			fmt.Println("Perfect")
		case 8,7: 
			fmt.Println("Awesome")
		default: 
			fmt.Println("Not Bad")
	}
}