package main

import "fmt"

func main() {
	energy := 10 //พลังงานเริ่มต้น (Max 10)
	gold := 0
	isWorking := true

	fmt.Println("--- Welcome to the Gold Mine ---")

	for isWorking {
		//show status now
		fmt.Printf("\n Energy: %d/10 | Gold: %d\n", energy, gold)
		fmt.Println("1. Mine Rock (Use 3 Energy)")
		fmt.Println("2. Sleep (Recover 5 Energy)")
		fmt.Println("3. Quit Job")
		fmt.Print("Choose action: ")

		var action int
		fmt.Scan(&action)

		switch action {
		case 1:
			//check strenge dig
			if energy >= 3 {
				energy = energy - 3
				gold = gold + 1
				fmt.Println(" Clang! You found a gold nugget!")
			} else {
				fmt.Println(" You are too tried! Go sleep.")
			}
		case 2:
			//Sleep for power
			energy = energy + 5
			fmt.Println("Zzz... You feel refreshed. ")

			//Logic เสริม: ห้าม Energy เกิน 10 (Max Stamina)
			if energy > 10 {
				energy = 10
				fmt.Println("Energy is full!)")
			}
		case 3:
			fmt.Println("Leaving the mine with total gold:", gold)
			isWorking = false //end loop
		default:
			fmt.Println(" Unknown action.")
		}
	}
}
