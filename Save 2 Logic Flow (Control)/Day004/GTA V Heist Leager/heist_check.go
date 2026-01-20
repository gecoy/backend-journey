package main

import "fmt"

func main() {
	var level int
	var money int

	fmt.Println("--- GTA V Heist Setup Check ---")

	fmt.Print("Enter your Level: ")
	fmt.Scanln(&level)

	fmt.Print("Enter your Money (&)")
	fmt.Scanln(&money)

	//USE && （AND）for 2 conditions
	//be Level >= 12 and Money >= 40000 ถึงจะผ่าน
	if level >= 12 && money >= 40000 {
		fmt.Println("Result: You can start the Heist! ")
	} else {
		//ถ้าขาดอย่างใดอย่างหนี่งหรือขาดทั้งคู่
		fmt.Println("Result: Not eligible yet. Go farm more! ")
	}
}
