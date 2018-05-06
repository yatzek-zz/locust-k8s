package main

import (
	"fmt"
	"log"
	"net/http"
)

func index_handler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("Index Request")
    if r.URL.Path != "/" {
        errorHandler(w, r, http.StatusNotFound)
        return
    }
	fmt.Fprintf(w, "Our example home page.")
}

func login_handler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("Login Request")
	fmt.Fprintf(w, "Login.")
}

func profile_handler(w http.ResponseWriter, r *http.Request) {
//	log.Printf("Profile Request")
	fmt.Fprintf(w, "Profile.")
}

func errorHandler(w http.ResponseWriter, r *http.Request, status int) {
    w.WriteHeader(status)
    if status == http.StatusNotFound {
        fmt.Fprint(w, "custom 404")
    }
}

func main() {
	http.HandleFunc("/", index_handler)
	http.HandleFunc("/login", login_handler)
	http.HandleFunc("/profile", profile_handler)
	log.Printf("Listening on port 8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
