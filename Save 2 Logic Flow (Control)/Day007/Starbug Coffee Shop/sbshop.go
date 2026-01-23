package main

import "fmt"

func main() {
	var size string
	var whip string
	var final int

	fmt.Println("--- Pick your DRINK!! ---")

	fmt.Println("Your cup size? (S/M/L): ")
	fmt.Scanln(&size)

	fmt.Println("You want whip? (yes/no): ")
	fmt.Scanln(&whip)

	if size == "S" {
		final = 40
	} else if size == "M" {
		final = 50
	} else if size == "L" {
		final = 60
	} else {
		fmt.Println("Error : โปรดสั่งสินค้าที่มีในเมนู")
		return // จบการทำงานทันที (ท่าไม้ตาย)
	}

	if whip == "yes" {
		final += 15 // ความหมายเหมือน final = final + 15 เป๊ะ!
	}

	fmt.Printf("Your price is: %d THB\n", final)

}
