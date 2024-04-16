package exampleGO

import (
	"fmt"
	"strings"
)

func ExampleMap() {
	// var maps map[string]int = map[string]int{}
	// fmt.Println(len(maps))
	// maps["Answer"] = 42
	// maps["Answer2"] = 103
	// fmt.Printf("%#v\n", maps)
	// for i := len(maps); i < 10; i++ {
	// }
	// for key, v := range maps {
	// 	fmt.Println(key, v)
	// }
	s := "If it looks like a duck swims like a duck and quacks like a duck then it probably is a duck"
	strings := make(map[string]int)
	strings = wordCounts(s)
	fmt.Println(strings)
}

func wordCounts(s string) map[string]int {
	maps := make(map[string]int)
	strings := strings.Fields(s)
	fmt.Println(strings)

	//add to map
	for _, v := range strings {
		maps[v]++
	}
	return maps
}
