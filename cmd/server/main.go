package main

import (
	"log"
	"net/http"

	"github.com/KenethSandoval/fofewi/internal/router"
	"github.com/KenethSandoval/fofewi/pkg/db"
	"github.com/KenethSandoval/fofewi/pkg/listening"
)

func main() {
	mux := http.NewServeMux()

	hs := &http.Server{
		Addr:    ":8080",
		Handler: mux,
	}

	listening.ListePrinServer(hs)
	router.InitRouter(mux)
	db.Connect()

	if err := hs.ListenAndServe(); err != nil {
		log.Fatalf("%v", err)
	}

}
