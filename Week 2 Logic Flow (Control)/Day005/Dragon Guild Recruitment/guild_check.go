package main

import "fmt"

func main() {
	var level int
	var hasLetter string //get a string input

	fmt.Println("--- Dragon Guild Recruitment ---")

	fmt.Print("Enter Level: ")
	fmt.Scanln(&level)

	fmt.Print("Do you have a recommendation letter? (yes/no): ")
	fmt.Scanln(&hasLetter)

	//use || (OR) for conditions
	//if level 30+ or have letter (อย่างใดอย่างหนึ่ง)
	if level >= 30 || hasLetter == "yes" {
		fmt.Println("Result: Welcome to the Guild! ")
	} else {
		fmt.Println("Result: You are not qualified. Get out!! ")
	}
}
