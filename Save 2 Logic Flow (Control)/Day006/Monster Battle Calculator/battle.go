package main

import "fmt"

func main() {
	var attack int
	var defense int
	var hp int
	var damage int

	fmt.Println("--- Battle Start! ---")

	//1.get Status
	fmt.Print("Enter Player Attack: ")
	fmt.Scanln(&attack)

	fmt.Print("Enter Monster Defense: ")
	fmt.Scanln(&defense)

	fmt.Print("Enter Monster HP: ")
	fmt.Scanln(&hp)

	fmt.Print("\n--- Result ---")

	//2.Calculation Damage (Check armor-piercing)
	if attack > defense {
		damage = attack - defense
		fmt.Printf("Critical Hit! You did %d damage. \n", damage)

		//3.cut hp monster (Nested If)
		hp = hp - damage

		if hp <= 0 {
			fmt.Println("Monster is DEFEATED! You obtain 100 EXP.")
		} else {
			fmt.Printf("Monster is angry! (HP remaining: %d)\n", hp)
		}
	} else {
		//cant do damage(Attack <= Defense)
		fmt.Println("Blocked! Your attack is too weak (0 Damage).")
	}
}
