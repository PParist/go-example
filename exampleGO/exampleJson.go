package exampleGO

import (
	"encoding/json"
	"fmt"
)

type Couse struct {
	Name        string `json:"name"`
	Level       string `json:"level"`
	Instuructor string `json:"instuructor"`
	Fullprice   int    `json:"fullprice"`
	View        int    `json:"view"`
}

type movie struct {
	Title       string   `json:"title"`
	Year        int      `json:"year"`
	Rating      float32  `json:"rating"`
	Genres      []string `json:"genres"`
	IsSuperHero bool     `json:"isSuperHero"`
}

func ExampleJson() {
	// _couse := Couse{
	// 	Name:        "GO",
	// 	Level:       "normal",
	// 	View:        7239,
	// 	Instuructor: "พิว",
	// 	Fullprice:   999,
	// }

	// _json, err := json.Marshal(_couse)
	// if err == nil {
	// 	fmt.Printf("%T and %s\n", _json, _json)
	// 	fmt.Printf("%v\n", string(_json))

	// 	_couse2 := new(Couse)
	// 	fmt.Printf("Befor Unamrshal : %T %#v\n", _couse2, _couse2)
	// 	err := json.Unmarshal(_json, _couse2)
	// 	if err == nil {
	// 		fmt.Printf("After Unmarshal : %#v\n", _couse2)
	// 	}
	// }

	//example
	body := `[
  {
    "imdbID": "tt4154796",
    "title": "Avengers: Endgame",
    "year": 2019,
    "rating": 8.4,
    "genres": ["Action", "Drama"],
    "isSuperHero": true
  },
  {
    "imdbID": "tt4154756",
    "title": "Avengers: Infinity War",
    "year": 2018,
    "rating": 8.4,
    "genres": ["Action", "Sci-Fi"],
    "isSuperHero": true
  }
]`
	moves := []movie{}
	err := json.Unmarshal([]byte(body), &moves)
	if err == nil {
		fmt.Printf("%#v\n", moves)
	}

}
