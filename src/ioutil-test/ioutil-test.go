package main

import (
	"fmt"
	"io/ioutil"
)

//Page data structure
type Page struct {
	Title string
	Body  []byte
}

//save feature
func (p *Page) save() error {
	filename := p.Title + ".txt"
	return ioutil.WriteFile(filename, p.Body, 0600)
}

//load feature
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func main() {
	p1 := &Page{Title: "Test Page", Body: []byte("This is a sample page.")}
	p1.save()
	p2, _ := loadPage("Test Page")
	fmt.Println(string(p2.Body))
}
