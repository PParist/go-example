package exampleGO

import "fmt"

func ExampleConstants() {

	type day int
	const (
		_ day = iota
		Sunday
		Monday
		Tuesday
		Wednesday
		Thursday
		Friday
		Saturday
	)

	fmt.Println("Sunday : ", Sunday)
	fmt.Println("Monday : ", Monday)
	fmt.Println("Tuesday : ", Tuesday)
	fmt.Println("Wednesday : ", Wednesday)
	fmt.Println("Thursday : ", Thursday)
	fmt.Println("Friday : ", Friday)
	fmt.Println("Saturday : ", Saturday)
	fmt.Printf("%T", Saturday)
}

func ExampleVariadic(value ...any) {
	for _, v := range value {
		fmt.Printf("Type is %T and raw value is %v\n", v, v)
	}
}
