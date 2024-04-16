package exampleGO

import (
	"fmt"
	"strconv"
)

func Convert() {
	_int := 256
	fmt.Printf("%T %#v\n", _int, _int)
	_float := float64(_int)
	fmt.Printf("%T %#v\n", _float, _float)

	_num := "2565566 aaaa"
	_result, err := strconv.Atoi(_num)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Printf("%T %#v\n", _result, _result)
	}

}
