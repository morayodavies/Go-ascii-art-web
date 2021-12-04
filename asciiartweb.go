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

	if r.URL.Path != "/" {
		handle404(w)
		return
	}
	fmt.Println("method:", r.Method) //get request method
	t, err := template.ParseFiles("html/template.gtpl")
	if err != nil {
		handleRequest(w)
		return
	}
	t.Execute(w, nil)
}

func asciiArt(w http.ResponseWriter, r *http.Request) {

	r.ParseForm()

	if len(r.Form["banner"]) == 0 {
		t, err := template.ParseFiles("html/404-banner-error.gtpl")
		if err != nil {
			handleRequest(w)
			return
		}
		w.WriteHeader(http.StatusNotFound)
		t.Execute(w, nil)
		return
	}

	words := strings.Split(strings.ReplaceAll(r.Form["word"][0], "\r\n", "\n"), "\n")

	for _, word := range words {
		PrintWord(w, word, r.Form["banner"][0])

	}
}

func PrintWord(w http.ResponseWriter, word string, banner string) {
	letters := []rune(word)

	//Validate input
	for _, letter := range letters {
		if letter < 32 || letter > 127 {
			t, err := template.ParseFiles("html/400-msg-error.gtpl")
			if err != nil {
				handleRequest(w)
				return
			}
			w.WriteHeader(http.StatusBadRequest)
			t.Execute(w, nil)
			return
		}
	}

	art := read_file("banners/" + banner)
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

func handleRequest(w http.ResponseWriter) {
	w.WriteHeader(http.StatusInternalServerError)
	t, err := template.ParseFiles("html/500-error.gtpl")
	if err != nil {
		log.Fatalf("Error happened in parsing file. Err: %s", err)
		return
	}
	t.Execute(w, nil)
}

func handle404(w http.ResponseWriter) {
	w.WriteHeader(http.StatusNotFound)
	t, err := template.ParseFiles("html/404-error.gtpl")
	if err != nil {
		log.Fatalf("Error happened in parsing file. Err: %s", err)
		return
	}
	t.Execute(w, nil)
}
