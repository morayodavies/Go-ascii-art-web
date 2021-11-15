package main

import (
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
)

func main() {
	http.HandleFunc("/", parseHtml) // setting router rule
	http.HandleFunc("/ascii-art", asciiArt)
	err := http.ListenAndServe(":9090", nil) // setting listening port
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}

func parseHtml(w http.ResponseWriter, r *http.Request) {
	fmt.Println("method:", r.Method) //get request method
	t, _ := template.ParseFiles("html/template.gtpl")
	t.Execute(w, nil)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if len(r.Form["banner"]) == 0 {
		t, _ := template.ParseFiles("html/banner-error.gtpl")
		t.Execute(w, nil)
		return
	}

	letters := []rune(r.Form["word"][0])

	//Validate input
	for _, letter := range letters {
		if letter < 32 || letter > 127 {
			t, _ := template.ParseFiles("html/msg-error.gtpl")
			t.Execute(w, nil)
			return
		}

	}

	art := read_file("banners/" + r.Form["banner"][0])
	printLetters(w, letters, art)
}

func read_file(s string) []string {
	// read in file
	content, err := ioutil.ReadFile(s)
	if err != nil {
		log.Fatal(err)
	}
	// replace newlines with space and random character and split using the random character
	text := strings.Replace(string(content), "\n", "%", -1)
	words := strings.FieldsFunc(text, func(r rune) bool { return strings.ContainsRune("%", r) })
	return words
}

func printLetters(w http.ResponseWriter, letters []rune, art []string) {
	for j := 0; j < 8; j++ {
		for i, letter := range letters {
			fmt.Fprintf(w, art[((int(letter)-32)*8)+j])
			if i == len(letters)-1 {
				fmt.Fprintf(w, "\n")
			}
		}
	}
}
