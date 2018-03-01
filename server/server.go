package main

import (
	"log"
	"net/http"
)

func main() {
	http.Handle("/submit", app1)
	http.Handle("/login", app2)
	log.Fatal(http.ListenAndServe("127.0.0.1:3000", nil))

}
