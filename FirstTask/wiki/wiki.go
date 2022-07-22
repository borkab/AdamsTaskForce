package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
)

//The Page struct describes
//how page data will be stored in memory.
type Page struct {
	Title string
	Body  []byte
}

/*
The Body element is a []byte
rather than string because that is the type expected
by the io libraries we will use
*/

func (p *Page) save() error { //This method will save the Page's Body to a text file.
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

/*
This is a method named save that takes as its receiver p,
 a pointer to Page . It takes no parameters,
and returns a value of type error.
*/
func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

/*
The function loadPage constructs the file name
from the title parameter, reads the file's contents
into a new variable body, and returns a pointer to a Page
 literal constructed with the proper title and body values.
*/

func main() {
	http.HandleFunc("/view/", viewHandler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}
