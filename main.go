package main

import (
	"fmt"
	"html/template"
	"log"
	"net/http"
)

type ContactFormData struct {
	Submitted bool
	Name      string
	Email     string
	Subject   string
	Message   string
	Error     string
}

var contactTmpl *template.Template

func init() {
	var err error
	contactTmpl, err = template.ParseFiles("./static/contact.html")
	if err != nil {
		log.Fatalf("Error parsing contact.html template: %v", err)
	}
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/hello" {
		http.NotFound(w, r)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Cannot take that request", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello from Shivam")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "GET" {
		http.Error(w, "Method not Allowed", http.StatusMethodNotAllowed)
		return
	}

	http.ServeFile(w, r, "./static/about.html")
}

func contactHandler(w http.ResponseWriter, r *http.Request) {
	data := ContactFormData{}

	if r.Method == "POST" {
		if err := r.ParseForm(); err != nil {
			data.Error = "Error parsing form. Please try again"
			w.WriteHeader(http.StatusBadRequest)
		} else {

			data.Name = r.FormValue("name")
			data.Email = r.FormValue("email")
			data.Subject = r.FormValue("subject")
			data.Message = r.FormValue("message")
			data.Submitted = true

			log.Printf("New Contact form submission: Name=%s, Email=%s", data.Name, data.Email)
		}
	}
	if err := contactTmpl.Execute(w, data); err != nil {
		log.Printf("Error executing contact template: %v", err)
		http.Error(w, "Internal Server Error", http.StatusInternalServerError)
	}
}

func main() {
	fmt.Println("Server starting on port 8080...")
	fileServer := http.FileServer(http.Dir("./static"))
	http.Handle("/", fileServer)
	http.HandleFunc("/hello", helloHandler)
	http.HandleFunc("/about", aboutHandler)
	http.HandleFunc("/contact", contactHandler)

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
