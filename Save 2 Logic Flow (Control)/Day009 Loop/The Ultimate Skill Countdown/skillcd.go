package main

import "fmt"

func main() {
	fmt.Println(" Warning: Ultimate Skill Charging... ")

	//loop ถอยหลัง
	//i := 5 -> start at 5
	//i > 0  -> ทำตราบใดที่ i ยังมากกว่า 0 (ถ้าเป็น 0 จะหยุด)
	//i--    ->  จบรอบแล้ว ลดค่าลงทีละ 1 (เช่น 5->4)
	for i := 5; i > 0; i-- {
		fmt.Printf("Casting in... %d\n", i)
	}

	//บรรทัดนี้จะทำงานเมื่อ Loop จบแล้ว (i กลายเป็น 0)
	fmt.Println(" Meteor Strike Deployed!!! ")
	fmt.Println("Enemy took 9999 Damage.")
}
