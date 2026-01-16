package main

import "fmt"

func main() {
	//1.pakard tua pare wait something
	var priceUSD float64
	var exchangeRate float64 = 34.5

	//2.ask user (prompt)
	fmt.Print("Enter game price in USD: ") //สั่งให้พิมพ์บรรทัดเดียวกัน

	//3. get something form keyboard(input)
	//program ja wait hai rao pim tua leg laew god ENTER
	fmt.Scanln(&priceUSD)

	//4.Calculation
	var priceTHB float64 = priceUSD * exchangeRate

	//5. output sadang pon
	fmt.Println("---Price Calculator---")
	fmt.Printf("Price in USD: $%.2f \n", priceUSD)
	fmt.Printf("Price in THB: %.2f Baht \n", priceTHB)

}
