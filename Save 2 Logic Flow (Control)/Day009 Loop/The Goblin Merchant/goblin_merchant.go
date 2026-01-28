package main

import "fmt"

func main() {
	//1.กำหนดค่าเริ่มต้น
	gold := 100
	shopOpen := true

	fmt.Println("--- Welcome to Goblin Shop ---")

	//2. Loop Shop
	for shopOpen {
		fmt.Printf("\n Your Gold: %d\n", gold)
		fmt.Println("1. Buy Sword (50 Gold)")
		fmt.Println("2. Buy Potion (20 Gold)")
		fmt.Println("3. Exit Shop")
		fmt.Print("Select item: ")

		var choice int
		fmt.Scan(&choice)

		fmt.Println("-----------------------------")

		switch choice {
		case 1:
			//Logic in Logic
			if gold >= 50 {
				gold = gold - 50 //ลดเงิน
				fmt.Println(" You bought a Sword! ")
			} else {
				//กรณีเงินไม่พอ
				fmt.Println("Not enough gold! Get out! ")
			}
		case 2:
			//Check medic
			if gold >= 20 {
				gold -= 20 //use -=
				fmt.Println("You bought a Potion! ")
			} else {
				fmt.Println(" You are too poor for this.")
			}
		case 3:
			fmt.Println("Come back when you have money.")
			shopOpen = false // close (End Loop)
		default:
			fmt.Println("?? What is that? We don't sell it.")
		}
	}

	fmt.Println("--- Shop Closed ---")
}
