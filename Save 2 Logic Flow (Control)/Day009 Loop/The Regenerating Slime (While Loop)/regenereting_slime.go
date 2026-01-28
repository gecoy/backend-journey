package main

import "fmt"

func main() {
	//set default
	slimeHP := 20
	turn := 1

	fmt.Println("--- Encountered a Wild Slime! (HP: 20) ---")

	//Loop แบบ "While" (ทำตราบใดที่เลือดSlimeยังมากกว่า 0)
	for slimeHP > 0 {
		//Vr Atk
		damage := 5

		//slime bleed
		//slimeHP -= damage มีค่าเท่ากับ slimeHP = slimeHP - damege
		slimeHP -= damage

		fmt.Printf("Turn %d: You hit for %d dmg! (Slime HP: %d)\n", turn, damage, slimeHP)

		//count more set
		turn++
	}

	fmt.Println("--- Slime Defeated! You win! ---")
}
