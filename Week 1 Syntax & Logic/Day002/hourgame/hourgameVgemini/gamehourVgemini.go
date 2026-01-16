package main

import "fmt"

func main() {
	// ใช้ CamelCase อ่านง่ายขึ้นเยอะ
	var gameOneHours float64
	var gameTwoHours float64

	fmt.Println("--- Game Hours Calculator ---")

	fmt.Print("Enter game one hours: ")
	fmt.Scanln(&gameOneHours) // ใช้ Scanln เพื่อรับค่าทีละบรรทัด

	fmt.Print("Enter game two hours: ")
	fmt.Scanln(&gameTwoHours)

	// วงเล็บ () ตรงบวกกันจริงๆ ไม่ต้องใส่ก็ได้ครับ Go ฉลาดพอ (แต่ใส่ก็ไม่ผิดนะ)
	var totalHours float64 = gameOneHours + gameTwoHours

	fmt.Printf("Total Game Hours: %.2f Hours \n", totalHours)
}
