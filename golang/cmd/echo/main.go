package main

import (
	"fmt"
	"github.com/fulls1z3/fullstack-monorepo/pkg/hello"
	"log"
	"net/http"
	"os"
)

func main() {
	var PORT string

	if PORT = os.Getenv("PORT"); PORT == "" {
		log.Fatal("PORT not defined")
	}

	var APP_NAME string

	if APP_NAME = os.Getenv("APP_NAME"); APP_NAME == "" {
		log.Fatal("APP_NAME not defined")
	}

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		msg := hello.GetHello(APP_NAME)
		fmt.Fprintf(w, "%s: %s\n", msg, r.URL.Path)
	})

	http.ListenAndServe(":" + PORT, nil)
}
