package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

//Payload to deliver
type Payload struct {
	Stuff Data
}

//Data structs
type Data struct {
	Fruit   Fruits
	Veggies Vegetables
}

//Fruits map a string to int
type Fruits map[string]int

//Vegetables map a string to int
type Vegetables map[string]int

func getJSON() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 15

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 4
	vegetables["Peppers"] = 0
	vegetables["Cheese"] = 48

	d := Data{fruits, vegetables}
	p := Payload{d}

	return json.MarshalIndent(p, "", " ")
}

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJSON()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:8080", nil)
}
