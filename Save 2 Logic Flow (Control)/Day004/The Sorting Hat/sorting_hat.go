package main

import "fmt"

func main() {
	var bravery int

	fmt.Println("--- The Sorting Hat ---")
	fmt.Print("Enter your bravery score (0-100): ")
	fmt.Scanln(&bravery)

	//Start Sorting
	if bravery >= 80 {
		fmt.Println("Result: Gryffindor! ")
	} else if bravery >= 60 {
		//if less than 80 ... but can go to 60?
		fmt.Println("Result: Ravenclaw! ")
	} else if bravery >= 40 {
		//if less than 60 .. but can go to 40?
		fmt.Println("Result: Hufflepuff! ")
	} else {
		// if mai kao condition
		fmt.Println("Result: Slytherin! ")
	}
}
