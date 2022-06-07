package main

import (
	"fmt"
	"net/http"
	"os"
)

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Running API V1"))
}

func main() {

	http.HandleFunc("/", rootHandler)

	err := http.ListenAndServe("localhost:11112", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
