package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Nama   string `json:"nama"`
	NIM    string `json:"NIM"`
	Alamat string `json:"Alamat"`
}

var data []Response

func Controller_Post(w http.ResponseWriter, r *http.Request) {
	//Header for return data MIME type
	w.Header().Set("Content-Type", "application/json")
	var req Response

	req.Nama = r.FormValue("Nama")
	req.NIM = r.FormValue("NIM")
	req.Alamat = r.FormValue("Alamat")

	data = append(data, req)
	extractData, err := json.Marshal(req)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprint(w, `{"status":"Error", "message":"Internal Server Error"}`)
		return
	}

	fmt.Fprint(w, string(extractData))
}