package main

import "fmt"

func main() {
	// 1. สร้าง Slice ว่างๆ (สังเกต [] ไม่มีตัวเลขข้างใน = ยืดหยุ่น)
	// ตอนนี้เหมือนเราถือใบรายชื่อเปล่าๆ
	var party []string

	// เช็คสถานะตอนแรก
	fmt.Printf("Start: %v (Size: %d)\n", party, len(party))

	// 2. เพิ่มตัวคุณเข้าไปคนแรก
	// party = ... คือการบอกว่า "ช่วยอัปเดตใบรายชื่อในมือฉันด้วยนะ"
	party = append(party, "G (Hero)")
	fmt.Println("... Leader joined!")

	// 3. เพิ่มเพื่อนใหม่ (ทีละคน)
	party = append(party, "Gandalf (Mage)")
	fmt.Println("... Mage joined!")

	// 4. เพิ่มเพื่อนใหม่อีก (เพิ่มทีละหลายคนก็ได้)
	party = append(party, "Legolas (Archer)", "Gimli (Fighter)")
	fmt.Println("... Archer & Fighter joined!")

	// 5. สรุปผล
	fmt.Println("\n-----------------------")
	fmt.Printf("Final Party Size: %d members\n", len(party))

	// วนลูปแนะนำตัวทีละคน
	for i, name := range party {
		fmt.Printf("Member #%d: %s\n", i+1, name)
	}
}
