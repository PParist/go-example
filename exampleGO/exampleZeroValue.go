package exampleGO

import "fmt"

func ExampleZeroValue() {
	var name_move string
	var year int
	var rating float64
	var category string
	var isSuper bool
	var r rune

	// fmt.Printf("เรื่อง: %s\n", name_move)
	// fmt.Printf("ปี: %d\n", year)
	// fmt.Printf("เรตติ่ง: %.1f\n", rating)
	// fmt.Printf("ประเภท: %s\n", category)
	// fmt.Printf("ซุปเปอร์ฮีโร่: %t\n", isSuper)
	// fmt.Printf("emoji = %c\n", r)

	fmt.Printf("เรื่อง: %#v Type : %T\n", name_move, name_move)
	fmt.Printf("ปี: %#v Type : %T\n", year, year)
	fmt.Printf("เรตติ่ง: %#v Type : %T\n", rating, rating)
	fmt.Printf("ประเภท: %#v Type : %T\n", category, category)
	fmt.Printf("ซุปเปอร์ฮีโร่: %#v Type : %T\n", isSuper, isSuper)
	fmt.Printf("r = %c Type : %T\n", r, r)
}
