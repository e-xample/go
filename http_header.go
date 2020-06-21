package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	contentType := r.Header.Get("Content-Type")
	fmt.Fprintln(w, contentType)

	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	v := struct {
		Msg string `json: "msg"`
	}{
		Msg: "hello",
	}

	if err := json.NewEncoder(w).Encode(v); err != nil {
		log.Fatal(err)
	}
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}
