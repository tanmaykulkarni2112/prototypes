package main

import (
	"net/http"

	"github.com/tanmaykulkarni2112/prototypes/golang-http/handlers"
)

func main() {
	// create the basic route for /home
	http.HandleFunc("/home", handlers.HomeHandler)

	http.HandleFunc("/api/home", handlers.HomeAPIHandler)
	// send json data as response

	// read from the file and send the content

	// read the post request and then update file and send updated content as response as json
	

	// create a middleware to ensure that user is autheneticated

	// create route to add cookie
	
	// login like route for cookie assigning
	http.ListenAndServe(":8080", nil)
}