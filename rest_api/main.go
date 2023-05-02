package main

import (
	"fmt"
	"net/http"
	"rest_api/controllers"
)


func main() {
	http.HandleFunc("/nama", controllers.Controller_Post)
	http.HandleFunc("/semuadata", controllers.Controller_Get )
    fmt.Println("Running on http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}