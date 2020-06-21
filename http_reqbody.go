package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Person struct {
	Name string `json: "name"`
	Age  int    `json: "age"`
}

func handler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	var p Person

	dec := json.NewDecoder(r.Body)
	if err := dec.Decode(&p); err != nil {
		log.Fatal(err)
	}
	fmt.Println(p)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
