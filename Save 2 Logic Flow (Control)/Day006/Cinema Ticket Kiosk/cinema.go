package main

import "fmt"

func main() {
	var basePrice int
	var age int
	var member string
	var finalPrice int

	fmt.Println("--- Cinema Ticket ---")

	fmt.Println("Your age: ")
	fmt.Scanln(&age)

	fmt.Println("Have member? (yes/no): ")
	fmt.Scanln(&member)

	if age > 60 {
		basePrice = 120
	} else if age >= 12 {
		basePrice = 200
	} else {
		basePrice = 100
	}

	if member == "yes" {
		finalPrice = basePrice - 20
	} else {
		finalPrice = basePrice
	}

	fmt.Printf("Final Price : %d THB\n", finalPrice)
}
