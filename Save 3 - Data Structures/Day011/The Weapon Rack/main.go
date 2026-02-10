package main

import "fmt"

func main() {
	//1.ประกาศ Array ขนาด 3 ช่อง เก็บข้อมูลแบบ String
	var weaponRack [3]string

	fmt.Println("--- Preparing Weapon Rack ---")

	// 2. ใส่ของเข้าไปในแต่ละช่อง (จำไว้! เริ่มที่ 0)
	weaponRack[0] = "Iron Sword"  //ช่องแรก
	weaponRack[1] = "Wooden Bow"  // ช่องสอง
	weaponRack[2] = "Magic Staff" // ช่องสาม

	// weaponRack[3] = "Dagger" <-- ถ้าบรรทัดนี้ทำงาน จะ Error ทันที! (เพราะมีแค่ 0-2)

	// 3. หยิบของออกมาดูทีละชิ้น
	fmt.Println("Slot 0 contains:", weaponRack[0])
	fmt.Println("Slot 1 contains:", weaponRack[1])
	fmt.Println("Slot 2 contains:", weaponRack[2])

	// 4. ดูของทั้งหมดในทีเดียว
	fmt.Println("\nFull Rack:", weaponRack)
}
