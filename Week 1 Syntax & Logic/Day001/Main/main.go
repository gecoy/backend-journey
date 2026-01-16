package main

import "fmt"

func main() {
	//1.pakard tua pare (Variables)
	//String = kep kor kuam
	var playerName string = "Stephen Curry"

	//int = kep jum nuan tem
	var shotsMade int = 8   //yinglong
	var shotsTotal int = 24 //yingtangmod

	//2.calculation
	//tong plang int to float64 (todsaniyom) ก่อน ไม่งั้นคอมจะงงถ้าเอาจำนวนเต็มมาหารกัน
	var accuracy float64 = (float64(shotsMade) / float64(shotsTotal)) * 100

	//3.print show result
	fmt.Println("Player:", playerName)
	//%.2f คือคำสั้่งให้showทศนิยม 2 ตำแหน่ง
	fmt.Printf("Shooting Accuracy: %.2f %% \n", accuracy)
}
