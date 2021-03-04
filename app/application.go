package app

import (
	"net/http"

	"github.com/gorilla/mux"
)

var (
	router = mux.NewRouter()
)

func StartApp() {
	mapUrls()

	srv := &http.Server{
		Handler: router,
		Addr:    "172.0.0.1:8000",
	}
	if err := srv.ListenAndServe(); err != nil {
		panic(err)
	}
}
