package main

import "fmt"

func main() {
	var TotalLoot float64
	var LesterFee float64
	var CrewMate float64

	fmt.Println("---Money we got---")

	fmt.Print("Total we loot: ")
	fmt.Scanln(&TotalLoot)

	fmt.Print("Lester fee pay: ")
	fmt.Scanln(&LesterFee)

	fmt.Print("Crew mate number: ")
	fmt.Scanln(&CrewMate)

	var Money float64 = (TotalLoot - LesterFee) / CrewMate

	fmt.Println("Money we made: ", Money)

}
