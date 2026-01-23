package main

import "fmt"

func main() {
	//set how many hit
	count := 5

	fmt.Println("--- Start Training ---")

	//start loop
	//i := 1    ---> เริ่มนับรอบที่ 1
	//i <= count ---> ทำตราบใดที่ i ยังน้อยกว่าหรือเท่ากับ 5
	//i++  ---> จบรอบแล้ว เพิ่มค่า i ไปอีก 1 (เช่น 1 -> 2)
	for i := 1; i <= count; i++ {
		fmt.Printf("Hit #%d: SMAAASH!! \n", i)
	}

	fmt.Println("--- Training Complete! ---")
	fmt.Println("You gained +50 EXP.")
}
