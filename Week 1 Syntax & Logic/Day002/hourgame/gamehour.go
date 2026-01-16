package main

import "fmt"

func main() {

	//mark
	var gameonehours float64
	var gametwohours float64

	fmt.Println("Game Hours Cal")

	fmt.Print("game one hours: ")
	fmt.Scan(&gameonehours)

	fmt.Print("game two hours: ")
	fmt.Scan(&gametwohours)

	//calculations
	var totalhours float64 = (gameonehours + gametwohours)

	fmt.Printf("Total Game Hours: %.2f Hours \n", totalhours)

}
