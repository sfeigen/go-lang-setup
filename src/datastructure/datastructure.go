package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Database struct {
	Node1 User
	Node2 Functions
	Node3 Access
}

type User struct {
	Name  Name
	Email Email
	Phone Phone
}

type Functions struct {
	Methods Methods
	Funcs   Funcs
	Queries Queries
}

type Access struct {
	Ground Low
	Mid    Medium
	Hight  Top
}

type Name map[int]string
type Email map[string]string
type Phone map[string]string

type Methods map[string]string
type Funcs map[string]string
type Queries map[string]string

type Low map[string]string
type Medium map[string]string
type Top map[string]string

type Payload struct {
	Load1 Database
}

func getJSON() ([]byte, error) {
	name := make(map[int]string)
	email := make(map[string]string)
	phone := make(map[string]string)

	methods := make(map[string]string)
	funcs := make(map[string]string)
	queries := make(map[string]string)

	low := make(map[string]string)
	mid := make(map[string]string)
	top := make(map[string]string)

	user := User{name, email, phone}
	function := Functions{methods, funcs, queries}
	access := Access{low, mid, top}
	db := Database{user, function, access}

	p := Payload{db}

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
