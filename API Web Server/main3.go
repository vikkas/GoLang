package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	err := http.ListenAndServe("localhost:11112", nil)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
