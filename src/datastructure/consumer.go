package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
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

func main() {
	url := "http://localhost:8080"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	var p Payload

	err = json.Unmarshal(body, &p)
	if err != nil {
		panic(err)
	}

	fmt.Println(p.Database.User)
}
