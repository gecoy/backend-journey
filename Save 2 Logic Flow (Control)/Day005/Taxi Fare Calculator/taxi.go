package main

import "fmt"

func main() {
	//1.set fare
	var totalFare int = 35
	var distance int
	var isTrafficJam string // yes/no

	fmt.Println("--- Taxi Fare Calculator ---")

	//get how far
	fmt.Print("Enter distance (km): ")
	fmt.Scanln(&distance)

	//2.calculation distance (Logic Flow)
	if distance <= 10 {
		//shrot distance: get 35 + (distance * 5)
		totalFare = totalFare + (distance * 5)
	} else {
		//long distance: get 35 + (distance * 10)
		totalFare = totalFare + (distance * 10)
	}
	//3.ask traffic jammed (Logic is separate from distance)
	fmt.Print("Is traffic jammed? (yes/no): ")
	fmt.Scanln(&isTrafficJam)

	if isTrafficJam == "yes" {
		//tech: + in own price
		totalFare = totalFare + 20
		fmt.Println("Traffic Charge: +20 Baht")
	}

	//4. summarize
	fmt.Printf("Total Fare: %d Baht \n", totalFare)
}
