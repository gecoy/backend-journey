package main

import "fmt"

func main() {
	// 1. สร้าง Slice ว่างๆ (สังเกต [] ไม่มีตัวเลข)
	var backpack []string

	fmt.Println("--- 🎒 Starting Adventure ---")
	fmt.Printf("Backpack status: %v (Size: %d)\n", backpack, len(backpack))

	// 2. เก็บของชิ้นแรก! (ใช้คำสั่ง append)
	// แปลว่า: เอา backpack เดิม มาต่อด้วย "Potion" แล้วเก็บกลับไปที่ backpack
	backpack = append(backpack, "Potion")
	fmt.Println("\nFound a Potion!")
	fmt.Printf("Backpack: %v (Size: %d)\n", backpack, len(backpack))

	// 3. เก็บของเพิ่มอีก 2 ชิ้น
	backpack = append(backpack, "Sword", "Map")
	fmt.Println("\nFound a Sword and a Map!")
	fmt.Printf("Backpack: %v (Size: %d)\n", backpack, len(backpack))

	// 4. วนลูปเช็คของ (ใช้ range เหมือน Array เป๊ะ!)
	fmt.Println("\n--- 🔍 Checking Inventory ---")
	for i, item := range backpack {
		fmt.Printf("Slot %d: %s\n", i+1, item)
	}
}
