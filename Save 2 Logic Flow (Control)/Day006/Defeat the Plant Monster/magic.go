package main

import "fmt"

func main() {
	var spell string    //Name of magic (fire, water, wind)
	var power int       //ATK BASE
	var finalDamage int //Last Damage deal (after cal magic)

	fmt.Println("--- Plant Monster appeared! ---")

	//1.get status name, power
	fmt.Print("Cast Spell (fire/water/wind): ")
	fmt.Scanln(&spell)

	fmt.Print("Enter Magic Power: ")
	fmt.Scanln(&power)

	//2. check condition magic (Logic Zone)
	if spell == "fire" {
		//fire is super effective *2 Damage
		finalDamage = power * 2
		fmt.Println("Result: Super Effective! Burning!! ")

	} else if spell == "water" {
		// water weakness damage == 0
		finalDamage = power * 0
		fmt.Println("Result: Not Effective... The plant grew bigger! ")

	} else {
		//another magic (wind, earth) normal damage
		finalDamage = power * 1
		fmt.Println("Result: Normal Hit.")
	}

	//3.summarize
	fmt.Printf("Total Damage dealt: %d \n", finalDamage)
}
