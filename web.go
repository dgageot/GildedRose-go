package main

import (
	"encoding/json"
	"net/http"
)

var inn = NewInn()

func init() {
	http.HandleFunc("/", handler)
}

func handler(w http.ResponseWriter, r *http.Request) {
	body, err := json.MarshalIndent(inn.Items, "", "  ")
	if err != nil {
		w.WriteHeader(500)
		return
	}

	w.Write(body)
}
