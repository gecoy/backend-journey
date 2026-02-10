package main

import "fmt"

func main() {
	// 1. ประกาศและใส่ค่าแบบย่อ (เรียกว่า Array Literal)
	// [3] คือขนาด, string คือชนิด, { ... } คือค่าข้างใน
	weapons := [3]string{"Iron Sword", "Wooden Bow", "Magic Staff"}

	fmt.Println("--- Scanning Inventory ---")
	// 2. ใช้ for range วนอ่านทุกช่อง
	// i คือ index (ลำดับช่อง), v คือ value (ชื่ออาวุธ)
	for i, v := range weapons {
		fmt.Printf("Slot #%d: [%s] is ready for battle!\n", i, v)
	}
	// Mentor Tip: ถ้าเราอยากนับจำนวนของทั้งหมดใน Array
	// ให้ใช้ฟังก์ชัน len() ย่อมาจาก length
	fmt.Printf("\nTotal items in rack: %d\n", len(weapons))
}
