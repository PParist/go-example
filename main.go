package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/PParist/go-example/exampleGoAPI"
)

func main() {
	//exampleFmt.ExampleFmt()
	// exampleZeroValue.ExampleZeroValue()
	// exampleIfElse.ExampleIfElse()
	// exampleSwitchCase.ExampleSwitchCase()
	// sum, num := examplefunction.ExampleFunction(2.4, 3.6)
	// fmt.Println(sum, num)
	// resaultCal := examplefunction.Ealculate
	// fmt.Printf("%T\n", resaultCal)
	// score := examplefunction.Emote(1.0)
	// fmt.Println(score)

	// fmt.Printf("%#v\n", name)

	// firstname := name[0]

	// fmt.Println(firstname)

	//exampleArray()
	//exampleLoop()
	//eampleSlice([]string{"Action", "Adventure", "Fantasy"})
	//eampleStruct()
	// c := couse{price: 990}
	// c.price = c.eampleMethod()
	// fmt.Printf("%#v\n", c)

	//m.eampleMethodInfo()
	//examplePointer()
	//exampleInterface()
	//exampleerror.ExampleError()
	//examplemap.ExampleMap()
	//exampleGO.Convert()
	//exampleGO.ExampleJson()
	//exampleGO.ExampleConstants()
	//exampleGO.ExampleVariadic("String", 10, 100.56, true)
	//exampleGO.ExampleDefer()
	http.HandleFunc("/movies", exampleGoAPI.MovieHandler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("Run on Port localhost:8080")
	}
}

func exampleArray() {
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	for i := 0; i < len(genres); i++ {
		fmt.Printf("genres[%d]: %s\n", i, genres[i])
	}
	genres[1] = "Sci-Fi"
	fmt.Printf("genres[1]: %s\n", genres[1])
}

func exampleLoop() {
	genres := [3]string{"Action", "Adventure", "Fantasy"}
	fmt.Printf("before for loop : %#v\n", genres)

	for index, _ := range genres {
		genres[index] = "Movie: " + genres[index]
	}
	fmt.Printf("after for loop : %#v\n", genres)

	for index, vaule := range genres {
		fmt.Printf("genres[%d]: %s\n", index, vaule)
	}
}

func eampleSlice(sl []string) {
	// skills := []string{"js", "go", "c#"}
	// fmt.Printf("Count : %d -- value: %#v\n", len(skills), skills)
	// skills = append(skills, "ts")
	// fmt.Printf("value : %#v\n", skills)
	// s1 := sl[:]
	// fmt.Printf("%#v", s1)

	xs := []float64{4, 5, 7, 8, 3, 8, 0}
	ys := []float64{7, 2, 10, 9, 7}
	vote := append(xs, ys...)
	fmt.Printf("votes : %v\n", vote)
	fmt.Printf("vote 5 - 8 : %v\n", vote[5:9])
}

func eampleStruct() {

	type movie struct {
		title       string
		year        int
		rating      float64
		genres      []string
		isSuperHero bool
	}

	m0 := movie{} //declared for zero value
	fmt.Printf("%#v\n", m0)

	m1 := movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}
	fmt.Printf("%#v\n", m1)

	m2 := movie{
		title:       "Infinity War",
		year:        2018,
		rating:      8.4,
		genres:      []string{"Action", "Sci-Fi"},
		isSuperHero: true,
	}
	fmt.Printf("%#v\n", m2)

	mvs := []movie{m1, m2}

	for _, v := range mvs {
		fmt.Printf("%#v\n", v)
	}
}

type movie struct {
	title       string
	year        int
	rating      float32
	votes       []float64
	genres      []string
	isSuperHero bool
}

var m = movie{
	title:       "Avengers: Endgame",
	year:        2019,
	rating:      8.4,
	genres:      []string{"Action", "Drama"},
	isSuperHero: true,
}

func (m movie) eampleMethodInfo() {

	fmt.Printf("%s (%d) - %.2f\n", m.title, m.year, m.rating)
	fmt.Println("Genders:")
	for _, v := range m.genres {
		fmt.Printf("\t%v\n", v)
	}
}

func changePrice(price *int) {

	fmt.Println("Address for Derefernce : ", price)
	if *price >= 599 {
		*price = *price - 599
		fmt.Println("On change price ", *price, &price)
	} else {
		fmt.Println("Price < 599")
	}

}

type couse struct {
	name, instuructor string
	price             float64
}

func examplePointer() {
	// price := 9999
	// addr := &price
	// fmt.Printf("Type of %T , Value of price =  %#v , Value of addr =  %#v\n", addr, price, addr)
	//price = 9999 - 400
	//changeValue(price)
	//fmt.Printf("Last : Value of price =  %#v , Value of addr =  %#v\n", price, &price)
	//_couse := &couse{name: "go", price: 9999}
	// _couse := new(couse)
	// fmt.Printf("addr is %v and type is %#v\n", &_couse, _couse)
	// discounted := _couse.discount()
	// fmt.Println("oder to user price is :", discounted.price)

	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	eg.addVote(8)
	fmt.Println("votes:", eg.votes)
}

func (m *movie) addVote(value float64) {
	m.votes = append(m.votes, value)
}
func changeValue(p int) int {
	fmt.Println("Change Value")
	p = p - 999
	addr := &p
	fmt.Printf("New : Type of %T , Value of price =  %#v , Value of addr =  %#v\n", addr, p, addr)
	return p
}

func (c couse) discount() couse {
	fmt.Println("Insert Couse", c.price)
	c.price = c.price - 999
	fmt.Println("Discounted is ", c.price)
	return c
}

func (m movie) showInfo() string {
	return m.title
}

type voter interface {
	addVote(float64)
}

func exampleInterface() {
	eg := &movie{
		title:       "Avengers: Endgame",
		year:        2019,
		rating:      8.4,
		votes:       []float64{7, 8, 9, 10, 6, 9, 9, 10, 8},
		genres:      []string{"Action", "Drama"},
		isSuperHero: true,
	}

	vote(eg, 8)
	fmt.Println("votes:", eg.votes)
}

func vote(v voter, _rating float64) {
	v.addVote(_rating)
}
