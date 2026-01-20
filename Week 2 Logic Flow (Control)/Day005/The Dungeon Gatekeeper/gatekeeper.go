package main

import "fmt"

func main() {
	var hasKey string
	var level int

	fmt.Println("--- The Dungeon Gate ---")

	//1. ask for a key first
	fmt.Println("Do you have the key? (yes/no): ")
	fmt.Scanln(&hasKey)

	//เริ่มตรวจสอบชั้นแรก
	if hasKey == "yes" {
		// --- โซนชั้นใน (door closed) ---
		fmt.Println("Door unlocked! But wait...")

		//2.ask for level (for players has a key)
		fmt.Print("Enter your Level: ")
		fmt.Scanln(&level)

		//check 2floor
		if level >= 20 {
			fmt.Println("Gatekeeper: Strong enough. Enter at your own risk! ")
		} else {
			fmt.Println("Gatekeeper: You are too weak! Go farm lvl 20 first. ")
		}
		//-------------------------------
	} else {
		//--- for the one who dont have a key
		fmt.Println("Nah bro door locked. You know you dont has key man.")
	}
}
