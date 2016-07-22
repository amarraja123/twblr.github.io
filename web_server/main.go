package main

import (
	"net/http"
	"fmt"
	"log"
)

func main()  {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request)  {
	fmt.Fprintln(w, "WelCome to first Go webServer")
}