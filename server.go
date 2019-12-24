package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type Page struct {
	Title string
	Body  []byte
}

func (pg *Page) save() error {
	filename := pg.Title + ".txt"
	return ioutil.WriteFile(filename, pg.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := ioutil.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{title, body}, nil
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[1:])
}

func pageSaveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	content := r.URL.Query()["content"]
	if len(content) == 0 {
		fmt.Fprintf(w, "Must provide page content in 'content' query param.")
		return
	}
	page := &Page{title, []byte(content[0])}
	err := page.save()
	if err != nil {
		fmt.Fprintf(w, "Page with title %s failed to save with err: %s", title, err)
		return
	}
	fmt.Fprintf(w, "Page with title %s saved.", page.Title)
}

func pageHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	page, err := loadPage(title)
	if err != nil {
		fmt.Fprintf(w, "<h1>%s</h1><div>Page failed to load with error: %s</div>", title, err)
	} else {
		fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", page.Title, page.Body)
	}
}

func main() {
	http.HandleFunc("/save/", pageSaveHandler)
	http.HandleFunc("/view/", pageHandler)
	http.HandleFunc("/", handler)
	log.Fatal(http.ListenAndServe(":8080", nil))
}
