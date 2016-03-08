package main

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/dgageot/gildedrose/inn"
)

func main() {
	inn := inn.NewInn()

	http.HandleFunc("/items", func(w http.ResponseWriter, r *http.Request) {
		body, err := json.MarshalIndent(inn.Items, "", "  ")
		if err != nil {
			w.WriteHeader(500)
			return
		}

		w.Write(body)
	})

	log.Fatal(http.ListenAndServe(":8080", nil))
}
