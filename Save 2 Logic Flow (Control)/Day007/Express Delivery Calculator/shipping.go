package main

import "fmt"

func main() {
	var weight float64
	var isRemote string
	var basePrice float64

	fmt.Println("--- Express Delivery!!! ---")

	fmt.Println("Your Shipping Weight: ")
	fmt.Scanln(&weight)

	if weight > 20 {
		fmt.Print("Overweight! We cannot ship.")
		return
	} else if weight <= 1 {
		basePrice = 40
	} else if weight <= 10 {
		basePrice = 70
	} else {
		basePrice = 100
	}

	fmt.Println("Remote Area? (yes/no): ")
	fmt.Scanln(&isRemote)

	if isRemote == "yes" {
		basePrice += 50
	}

	fmt.Printf("Your Shipping Price is: %.2f THB \n", basePrice)

}
