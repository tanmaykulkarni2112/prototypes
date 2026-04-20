package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"os"
)

// for /home route , send a welcome message as response
var HomeHandler = func(w http.ResponseWriter, _ *http.Request) {
	words , err := w.Write([]byte("Welcome to Home page"))
	if err != nil {
		panic(err)
	}
	fmt.Println("Number of words ", words )
}

// for /api/home route, send json response

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

// handling /api/readfile route, read from text file and send json response

type FileResponse struct {
    FileName    string `json:"fileName"`    // ← exported
    FileContent string `json:"fileContent"` // ← exported
}

var ReadFileHandler = func (w http.ResponseWriter, r *http.Request) {
	// In Go, relative paths are resolved from the process’s working directory,
	// not from the file where the function is written.

	// So when you run your program (for example go run main.go or executing the built binary),
	// the current working directory is what determines how a relative path is interpreted.
	path := "./docs/textfile.txt"
	data, err := os.ReadFile(path) // returns byte[]
	if err != nil {
		http.Error(w, "File not found", http.StatusNotFound)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	response := FileResponse{
		FileName: path,
		FileContent: string(data),
	}

	if err := json.NewEncoder(w).Encode(response); err != nil {
		http.Error(w, "Failed to send response", http.StatusInternalServerError)
		return 
	}

}