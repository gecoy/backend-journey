package main

import "fmt"

func main() {
	// 1. กำหนดรหัสลับ (คุณตั้งเป็นเลขอื่นก็ได้นะ)
	secretNumber := 72

	fmt.Println("--- 🗝️ The Cursed Lock ---")
	fmt.Println("Guess the number between 1-100 to open the chest!")

	// 2. เริ่ม Loop อมตะ (ไม่มีเงื่อนไขหยุด จนกว่าจะสั่ง break)
	for {
		fmt.Print("Enter your guess: ")
		var guess int
		fmt.Scan(&guess)

		// 3. เช็คเงื่อนไข
		if guess < secretNumber {
			fmt.Println("🔼 Too Low! (ลองเลขที่สูงกว่านี้)")
		} else if guess > secretNumber {
			fmt.Println("🔽 Too High! (ลองเลขที่ต่ำกว่านี้)")
		} else {
			// กรณีที่ทายถูก (guess == secretNumber)
			fmt.Println("🎉 CORRECT! The chest opens.")
			fmt.Println("You found: 🏆 The Golden Keyboard!")

			// 4. สั่งทำลาย Loop ทันที
			break
		}
	}

	fmt.Println("--- Game Over ---")
}
