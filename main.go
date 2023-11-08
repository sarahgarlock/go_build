package main

import(
	"fmt" // formatting output 
	"log" // logging errors
	"net/http" // building an HTTP server
)

/* This function is an HTTP request handler for the path "/form." 
It handles POST requests, attempts to parse form data, and extracts 
the values of the "name" and "address" fields from the form. 
It then writes a response with the received data. */

func formHandler(w http.ResponseWriter, r *http.Request) {
	// Attempt to parse form data from the request.
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "ParseForm() err: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful\n")
	// Get the values of form fields named "name" and "address."
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}

/* This function is another HTTP request handler, 
but it's for the path "/hello." 
It checks if the request path is "/hello" 
and if the request method is GET. 
If these conditions are met, it responds with "Hello!".*/

func helloHandler(w http.ResponseWriter, r *http.Request) {
	// Check if the request path is "/hello."
	if r.URL.Path != "/hello" {
		http.Error(w, "404 not found.", http.StatusNotFound)
		return
	}
	// Check if the request method is GET.
	if r.Method != "GET" {
		http.Error(w, "Method is not supported.", http.StatusNotFound)
		return
	}
	fmt.Fprintf(w, "Hello!")
}

/*The main function sets up a file server to serve static files 
from the "static" directory, registers the 
formHandler and helloHandler functions to handle 
specific paths, and starts the HTTP server on port 8080. 
Any incoming requests will be handled by the registered request handlers.*/
func main(){
	// Serve static files from the "static" directory.
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	// Register request handlers for "/form" and "/hello."
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	// Start the HTTP server on port 8080.
	fmt.Printf("Starting server at port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err) // log package
	}
}