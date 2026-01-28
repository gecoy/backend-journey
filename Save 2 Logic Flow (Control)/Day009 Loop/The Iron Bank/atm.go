package main

import "fmt"

func main() {
	//เงินเริ่มต้นในบัญชี
	balance := 1000
	isUsingATM := true

	fmt.Println("--- Welcome to Iron Bank ---")

	for isUsingATM {
		fmt.Printf("\n Current Balance: %d THB\n", balance)
		fmt.Println("1. Deposit (ฝากเงิน)")
		fmt.Println("2. Withdraw (ถอนเงิน)")
		fmt.Println("3. Exit (ออก)")
		fmt.Print("Choose option: ")

		var choice int
		fmt.Scan(&choice)

		switch choice {
		case 1:
			fmt.Print("Enter amount to deposit: ")
			var amount int
			fmt.Scan(&amount)
			//ฝากเงิน = เอาเงินเก่า + เงินใหม่
			balance = balance + amount
			fmt.Println(" Deposit Success!")

		case 2:
			fmt.Print("Enter amount to withdraw: ")
			var amount int
			fmt.Scan(&amount)

			//Logic: เช็คว่าเงินพอให้ถอนไหม?
			if amount > balance {
				fmt.Println(" Error: Insufficient funds! (เงินไม่พอ)")
			} else {
				balance = balance - amount
				fmt.Println(" Withdraw Success! Here is your money.")
			}
		case 3:
			fmt.Println(" Thank you for using Iron Bank service.")
			isUsingATM = false // end Loop

		default:
			fmt.Println("Wrong option, please try again.")
		}
	}
}
