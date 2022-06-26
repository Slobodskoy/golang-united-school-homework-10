package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strconv"

	"github.com/gorilla/mux"
)

/**
Please note Start functions is a placeholder for you to start your own solution.
Feel free to drop gorilla.mux if you want and use any other solution available.

main function reads host/port from env just for an example, flavor it following your taste
*/

// Start /** Starts the web server listener on given host and port.
func Start(host string, port int) {
	router := mux.NewRouter()
	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
	})

	router.HandleFunc("/name/{PARAM}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)
		param := vars["PARAM"]
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hello, %v!", param)
	}).Methods("GET")

	router.HandleFunc("/bad", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusInternalServerError)
	}).Methods("GET")

	router.HandleFunc("/data", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(http.StatusOK)
		param := r.FormValue("PARAM")
		fmt.Fprintf(w, "I got message:\n%s", param)
	}).Methods("POST")

	router.HandleFunc("/headers", func(w http.ResponseWriter, r *http.Request) {
		s_a, s_b := r.Header.Get("a"), r.Header.Get("b")
		a, _:= strconv.Atoi(s_a)
		b, _ := strconv.Atoi(s_b)
		sum := a + b
		w.Header()["a+b"] = []string {strconv.Itoa(sum)}
		w.WriteHeader(http.StatusOK)
	}).Methods("POST")

	log.Println(fmt.Printf("Starting API server on %s:%d\n", host, port))
	if err := http.ListenAndServe(fmt.Sprintf("%s:%d", host, port), router); err != nil {
		log.Fatal(err)
	}
}

//main /** starts program, gets HOST:PORT param and calls Start func.
func main() {
	host := os.Getenv("HOST")
	port, err := strconv.Atoi(os.Getenv("PORT"))
	if err != nil {
		port = 8081
	}
	Start(host, port)
}
