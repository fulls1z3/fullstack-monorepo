package main

import (
	"fmt"
	"github.com/fulls1z3/fullstack-monorepo/pkg/hello"
	"net/http"
	"os"
)

func main() {
	var PORT string

	if PORT = os.Getenv("PORT"); PORT == "" {
		PORT = "3001"
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		hello := hello.GetHello()
		fmt.Fprintf(w, "%s: %s\n", hello, r.URL.Path)
	})

	http.ListenAndServe(":"+PORT, nil)
}
