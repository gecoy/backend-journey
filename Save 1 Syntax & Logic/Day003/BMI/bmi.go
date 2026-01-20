package main

import "fmt"

func main() {

	var Weight float64
	var Height float64

	fmt.Println("---Bmi NBA Players---")

	fmt.Print("Enter Weight: ")
	fmt.Scanln(&Weight)

	fmt.Print("Enter Height: ")
	fmt.Scanln(&Height)

	var BMI float64 = Weight / (Height * Height)

	fmt.Println("Player BMI is: ", BMI)
}
