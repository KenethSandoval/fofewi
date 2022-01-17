package router

import (
	"fmt"
	"net/http"

	"github.com/KenethSandoval/fofewi/internal/router/users"
)

type apiHandler struct{}

func (apiHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
}

func InitRouter(mux *http.ServeMux) {
	mux.HandleFunc("/", homePage)
	mux.HandleFunc("/user", users.GetAll)
}

func homePage(w http.ResponseWriter, r *http.Request) {
	// The "/" pattern matches everything, so we need to check
	// that we're at the root here.
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	fmt.Fprintf(w, "Welcome to the home page!")
}
