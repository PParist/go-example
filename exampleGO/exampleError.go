package exampleGO

import (
	"errors"
	"fmt"
)

func ExampleError() {
	sum, err := Divide(1, 0)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(sum)
	}
}
func Divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("cannot divide by zero")
	}

	resault := a / b

	return resault, nil
}
