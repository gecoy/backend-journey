package main

import "fmt"

func main() {
	var gold int
	var choice int
	var remainingGold int

	fmt.Println("--- Welcome to Item Shop ---")

	//1. get money we have
	fmt.Print("How much gold do you have?: ")
	fmt.Scanln(&gold)

	//2.Show menu
	fmt.Println("\nItems for sale: ")
	fmt.Println("1. Healing Potion (50 G)")
	fmt.Println("2. Magic Sword    (200 G)")
	fmt.Println("3. Dragon Shield  (500 G)")

	fmt.Print("Choose item (1-3): ")
	fmt.Scanln(&choice)

	fmt.Println("\n--- Transaction Result ---")

	//3.Start check condition (Nested If or Logic normal)
	if choice == 1 {
		// wants Potion (50 G)
		if gold >= 50 {
			remainingGold = gold - 50 // cut the money
			fmt.Println("You bought: Healing Potion! ")
			fmt.Println("Remaining Gold:", remainingGold)
		} else {
			fmt.Println("Not enough gold for Potion! ")
		}
	} else if choice == 2 {
		//Need Sword (200 G)
		if gold >= 200 {
			remainingGold = gold - 200
			fmt.Println("You bought: Magic Sword! ")
		} else {
			fmt.Println("Not enough gold for Sword! ")
		}
	} else if choice == 3 {
		//need Shield (500 G)
		if gold >= 500 {
			remainingGold = gold - 500
			fmt.Println("You bought: Dragon Shield! ")
			fmt.Println("Remaining Gold:", remainingGold)
		} else {
			fmt.Println("Not enough gold for Shield! ")
		}
	} else {
		//use another num
		fmt.Println("Item not found. Please choose 1-3 only.")
	}
}
