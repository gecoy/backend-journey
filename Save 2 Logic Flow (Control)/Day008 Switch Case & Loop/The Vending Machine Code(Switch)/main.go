package main

import "fmt"

func main() {
	var choice int

	//show menu
	fmt.Println("=== Vending Machine of Fate ===")
	fmt.Println("1. Blue Potion (Mana)")
	fmt.Println("2. Red Potion (Health)")
	fmt.Println("3. Green Potion (Poison Cure)")
	fmt.Println("4. Exit")
	fmt.Print("Choose your item (1-4): ")

	fmt.Scan(&choice)

	fmt.Println("\n--- Processing ---")

	switch choice {
	case 1:
		fmt.Println("You got a Blue Potion! (Mana +50)")
	case 2:
		fmt.Println("You got a Red Potion! (HP +100)")
	case 3:
		fmt.Println("You got a Green Potion! (Status Clear)")
	case 4:
		fmt.Println("Good Luck on your journey!")
	default:
		fmt.Println("Error: Invalid selection. Coin Returned.")
	}
}
