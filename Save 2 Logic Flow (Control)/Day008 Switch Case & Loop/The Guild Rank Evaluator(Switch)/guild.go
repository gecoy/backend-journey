package main

import "fmt"

func main() {
	var rank string

	fmt.Println("Enter your Hunter Rank (S, A, B, C): ")
	fmt.Scan(&rank)

	switch rank {
	case "S":
		fmt.Println("Level: God Tier (Access to Dragon Zone)")
	case "A":
		fmt.Println("Lecel: Elite Hunter (Access to High Danger Zone)")
	case "B":
		fmt.Println("Level: Experienced (Access to Normal Zone)")
	case "C":
		fmt.Println("Level: Rookie (Access to Training Ground)")
	default:
		fmt.Println("Invalid Rank. Access Denied.")
	}

}
