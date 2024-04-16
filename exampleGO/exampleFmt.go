package exampleGO

import "fmt"

func ExampleFmt() {
	name_move := "Avengers: Endgame"
	year := 2019
	rating := 8.4
	category := "Sci-Fi"
	isSuper := true
	var r rune = '\U0001F60D'

	fmt.Printf("เรื่อง: %s\n", name_move)
	fmt.Printf("ปี: %d\n", year)
	fmt.Printf("เรตติ่ง: %.1f\n", rating)
	fmt.Printf("ประเภท: %s\n", category)
	fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuper)
	fmt.Printf("emoji = %c\n", r)

	// fmt.Printf("เรื่อง: %#v Type : %T\n", name_move, name_move)
	// fmt.Printf("ปี: %#v Type : %T\n", year, year)
	// fmt.Printf("เรตติ่ง: %#v Type : %T\n", rating, rating)
	// fmt.Printf("ประเภท: %#v Type : %T\n", category, category)
	// fmt.Printf("ซุปเปอร์ฮีโร่: %#v Type : %T\n", isSuper, isSuper)
	// fmt.Printf("r = %c Type : %T\n", r, r)
}
