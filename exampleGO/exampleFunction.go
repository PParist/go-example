package exampleGO

import "fmt"

func ExampleFunction(x, y float64) (float64, string) {
	fmt.Printf("%.1f\n", x)
	return x, "hello"
}

func ExampleNakedReturn(sum int) (x int, y int) {
	x = sum / 2
	y = sum - x
	return
}

var _string string = "Function"
var Ealculate func(int, int) (int, string) = func(x, y int) (int, string) {
	x += 2
	y *= 10
	return x + y, _string
}

func Emote(ratings float64) string {
	switch {
	case ratings < 5.0:
		return "Disappointed 😞"
	case ratings >= 5.0 && ratings < 7.0:
		return ("Normal 😐")
	case ratings >= 7.0 && ratings < 10:
		return ("Good 🥰")
	default:
		return ("🤷🤷🤷🤷")
	}
}
