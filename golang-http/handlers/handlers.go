package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
)

var HomeHandler = func(w http.ResponseWriter, _ *http.Request) {
	words , err := w.Write([]byte("Welcome to Home page"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of words ", words )
}

type HomeResponse struct {
	Message string `json:"message"`
}

var HomeAPIHandler = func (w http.ResponseWriter, _ *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	response := HomeResponse{
		Message: "Welcome to api/home page",
	}
	json.NewEncoder(w).Encode(response)
}