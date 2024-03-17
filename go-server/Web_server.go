package main

import (
	"fmt"
	"log"
	"net/http"
)

func main(){
	fileServer := http.FileServer(http.Dir("./static")) //check static folder
	http.Handle("/", fileServer) 
	http.HandleFunc("/form", formHandler) // form.html
	http.HandleFunc("/hello", helloHandler) //print hello to the screen

	fmt.Printf(("Starting server at port 8080\n"))

	if err := http.ListenAndServe(":8080", nil); err != nil{
		log.Fatal(err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request){ //take request and response here * is pointer
	if r.URL.Path != "/hello"{
		http.Error(w, "404 not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET"{
		http.Error(w, "Method Not Supported", http.StatusNotFound)
	}
	fmt.Fprintf(w, "Hello")
}

func formHandler(w http.ResponseWriter, r *http.Request){
	if err := r.ParseForm(); err != nil{
		fmt.Fprintf(w, "Parseform() err : %v", err)
		return
	}
	fmt.Fprintf(w, "POST request succesfull")
	name := r.FormValue("name")
	address := r.FormValue("address")
	fmt.Fprintf(w, "Name = %s\n", name)
	fmt.Fprintf(w, "Address = %s\n", address)
}