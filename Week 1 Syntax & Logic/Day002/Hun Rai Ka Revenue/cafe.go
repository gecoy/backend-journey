package main

import "fmt"

func main() {
	//1.make price of coffee(use float cuz have todsaniyom)
	var price float64 = 65.50

	//2.get a cups (use int cuz cup cant have todsanitom)
	var quantity int

	fmt.Println("--- Hun Rai Ka Daily Sales ---")
	fmt.Print("How many cups sold today?: ")
	fmt.Scanln(&quantity) //dont forget & important cuz & want where is variable

	//3.Calculations
	//we have to change quantity(int) go to (float64) first before multiply price
	var totalRevenue float64 = price * float64(quantity)

	//4.Show somethings we made
	fmt.Printf("Total Revenue: %.2f Baht \n", totalRevenue)
}
