package main

import "fmt"

func main() {
	//1.pakard tua pare (use int cuz ball score dont have todsaniyom)
	var threePointers int
	var twoPointers int
	var freeThrows int

	//2.get something from keyboard
	fmt.Println("---Steph Curry Score Calculator ---")

	fmt.Print("Enter 3-Pointers made:")
	fmt.Scanln(&threePointers) //dont forget & important

	fmt.Print("Enter 2-Pointers made:")
	fmt.Scanln(&twoPointers)

	fmt.Print("Enter 1-Pointers made:")
	fmt.Scanln(&freeThrows)

	//3.Calculations (Logics)
	//สูตร: (3Pointers * 3) + (2Pointers * 2) + (Freethrows * 1)
	var totalScore int = (threePointers * 3) + (twoPointers * 2) + (freeThrows * 1)

	//4.Show Somethings
	fmt.Printf("Total Score: %d points \n", totalScore)
	// %d use to show integer or decimal

}
