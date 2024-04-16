package exampleGO

import "fmt"

func ExampleIfElse() {
	rating := 8.4
	sadEmote := "😔"
	normalEmote := "😐"
	happyEmote := "🥰"
	ingoreEmote := "🤷🏻‍♂️🤷🏻‍♂️🤷🏻‍♂️🤷🏻‍♂️"

	if rating < 5.0 {
		fmt.Println("Disappointed", sadEmote)
	} else if rating >= 5.0 && rating <= 7.0 {
		fmt.Println("Normal", normalEmote)
	} else if rating >= 7.0 && rating < 10.0 {
		fmt.Println("Good", happyEmote)
	} else {
		fmt.Println(ingoreEmote)
	}
}
