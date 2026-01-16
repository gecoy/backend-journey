package main

import "fmt"

func main() {
	var score int

	fmt.Println("--- Potinos Exam Result ---")
	fmt.Print("Enter your score: ")
	fmt.Scanln(&score)

	//first time use logic flow (If-Else)
	if score >= 50 {
		//if condition true (true) ja play peek ka
		fmt.Println("Result: PASS! You are a wizard! @_@")
	} else {
		//if condition false (False) ja play this peek ka tan
		fmt.Println("Result: FAIL! Go home muggle. TT")
	}

}
