package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
        w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
        w.Header().Set("Access-Control-Allow-Headers", "Content-Type, Authorization")
		w.WriteHeader(http.StatusOK)
		fmt.Fprintf(w, "Hallo, Elastic Beanstalk!")
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe("127.0.0.1:5000", nil)
}
