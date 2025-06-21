package server

import (
	"fmt"
	"net/http"
)

func Server(){
	mux := http.NewServeMux()

	// Endpoints
	mux.HandleFunc("/", handleRoot) // Root
	
	// Start Server
	fmt.Println("Server is listening to :8080")
	http.ListenAndServe(":8080", mux)
}

// Root Endpoint
func handleRoot(
	w http.ResponseWriter, 
	r *http.Request,
	){
	fmt.Fprintf(w, "Hello World!")
}