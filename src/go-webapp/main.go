package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	port, ok := os.LookupEnv("PORT")
	if !ok {
		port = "8080"
	}

	hostname, _ := os.Hostname()
	http.HandleFunc("/", func(w http.ResponseWriter, _ *http.Request) {
		fmt.Fprintf(w, "[%s] Success! Go webapp running\n", hostname)
	})

	http.ListenAndServe(":"+port, nil)
}
