package exampleGO

import "fmt"

func ExampleDefer() {
	//TODO :ใช้ Defer เพื่อทำอะไรซักอย่างหลังจบ func

	// defer fmt.Println("one")
	// defer fmt.Println("two")
	// defer fmt.Println("three")

	//fmt.Println("zero")
	fmt.Println("Counting")
	for i := 0; i < 10; i++ {
		fmt.Println("i = ", i)
		defer func(n int) {
			fmt.Println(n)
		}(i)
	}

}
