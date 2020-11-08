package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)
	r.HandleFunc("/{key}", LogHandler)
	http.Handle("/", r)
	http.ListenAndServe(":8080", nil)
}
func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "ShellyMicroWeb Action Logger")
}

func LogHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	req := getIPAddress(r)
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Key: %v\n", vars["key"])
	f, err := os.OpenFile("webserver.log", os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("error opening file: %v", err)
	}
	defer f.Close()

	log.SetOutput(f)

	log.Printf("%s - %v\n", req, vars["key"])
}

func getIPAddress(r *http.Request) string {

	ip := strings.SplitAfter(r.RemoteAddr, ":")
	s := strings.TrimSuffix(ip[0], ":")
	return s
}
