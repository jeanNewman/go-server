package main

import(
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err!= nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}
	fmt.Fprintf(w, "post request suscessfully")
	name := r.PostForm.Get("name")
	address := r.PostForm.Get("address")
	fmt.Fprintf(w, "name: %s, address: %s", name, address)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path!= "/hello" {
		http.Error(w, "404 not found", http.StatusNotFound)
		
        return
	}
	if r.Method!= "GET" {
		http.Error(w, "405 method not allowed", http.StatusMethodNotAllowed)
        
        return
	}
	fmt.Fprintf(w, "Hello, %q", r.URL.Path)
}


func main() {
	fileserver := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileserver)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Listening on port 8080\n")
	if err := http.ListenAndServe(":8080", nil); err!= nil {
		log.Fatal(err)
	}
}
