package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/claudiovictors/http-json/foo"
)

func main() {
	app := http.NewServeMux()
	app.HandleFunc("GET /posts", foo.Index)
	app.HandleFunc("GET /posts/{id}", foo.Show)
	app.HandleFunc("POST /posts", foo.Store)

	fmt.Println("Server runing in port 8080")
	log.Fatal(http.ListenAndServe(":8080", app))
}
