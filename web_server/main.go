package main

import (
	"net/http"
	"fmt"
	"log"
	"regexp"
)

func main()  {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {
	re := regexp.MustCompile("^(.+)@golang.org$")
	path := r.URL.Path
	match := re.FindAllStringSubmatch(path, -1)
	if match != nil{
		fmt.Fprintf(w, "Hello %s\n", match[0][1])
		return
	}
	fmt.Fprintln(w, "I do not recongize you ", path)
}