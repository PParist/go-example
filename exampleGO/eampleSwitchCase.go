package exampleGO

import "fmt"

func ExampleSwitchCase() {

	rating := 8.4
	sadEmote := "😔"
	normalEmote := "😐"
	happyEmote := "🥰"
	ingoreEmote := "🤷🏻‍♂️🤷🏻‍♂️🤷🏻‍♂️🤷🏻‍♂️"

	switch {
	case rating < 5.0:
		fmt.Println("Disappointed", sadEmote)
	case rating >= 5.0 && rating < 7.0:
		fmt.Println("Normal", normalEmote)
	case rating >= 7.0 && rating < 10.0:
		fmt.Println("Good", happyEmote)
	default:
		fmt.Println(ingoreEmote)
	}
}
