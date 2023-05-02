package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func Controller_Get(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	extractData, err := json.Marshal(data)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status":"Error", "message":"Internal Server Error"}`)
		return
	}
	fmt.Fprint(w, string(extractData))
}