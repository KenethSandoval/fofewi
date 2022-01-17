package main

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/fofewi/pkg/listening"
)

func main() {
	mux := http.NewServeMux()

	hs := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	listening.ListePrinServer(hs)

	if err := hs.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}

}
