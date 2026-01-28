package main

import "fmt"

func main() {
	hp := 100
	isRunning := true //virable important

	//Loop ตราบใดที่ isGameOver ยังเป็น false
	for isRunning {

		fmt.Printf("\n--- Status HP: %d ---\n", hp)
		fmt.Println("1. Attack")
		fmt.Println("2. Potion")
		fmt.Println("3. Exit")
		fmt.Print("Choose: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			hp = hp - 30
		case 2:
			hp = hp + 20
		case 3:
			fmt.Println("Bye Bye!")
			isRunning = false
		default:
			fmt.Println("Error: Wrong input")
		}
		if hp <= 0 {
			fmt.Println(" You died! (HP is 0 or less)")
			isRunning = false //สั่งจบเกมทันที!
		}

	}

	fmt.Println("--- Game Over ---")

}
