package exampleGoAPI

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

type Movie struct {
	ImdbID      string  `json:"imdbID"`
	Title       string  `json:"title"`
	Year        int     `json:"year"`
	Rating      float32 `json:"rating"`
	IsSuperHero bool    `json:"isSuperHero"`
}

var Movies []Movie

// TODO:this func is recive req and return res
func MovieHandler(w http.ResponseWriter, req *http.Request) {
	//check is get post push del... any
	method := req.Method
	switch {
	case method == "POST":
		body, err := ioutil.ReadAll(req.Body)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
			w.Write([]byte(err.Error()))
			//fmt.Fprintf(w, "error : %v\n", err)
		} else {
			//fmt.Fprintf(w, "body is %s\n", string(body))
			m := Movie{}
			err := json.Unmarshal([]byte(body), &m)
			if err == nil {
				w.WriteHeader(http.StatusOK)
				Movies = append(Movies, m)
				fmt.Fprintf(w, "hello %s create movies success\n", method)
			} else {
				w.WriteHeader(http.StatusBadRequest)
				//fmt.Printf("error : %v", err)
				w.Write([]byte(err.Error()))
			}
		}
	case method == "GET":
		_json, err := json.Marshal(Movies)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintf(w, "error %s", err)
		} else {
			w.WriteHeader(http.StatusOK)
			//fmt.Fprintf(w, "hello %s movies\n data : %v", method, string(_json))
			w.Header().Set("content-type", "application/json;charset=utf-8")
			w.Write(_json)
		}
	default:
		fmt.Fprintf(w, "hello %s movies\n", method)
	}
}
