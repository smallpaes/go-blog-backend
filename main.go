package main

import (
	"net/http"
	"time"

	"github.com/smallpaes/go-blog-backend/internal/routers"
)

func main() {
	router := routers.NewRouter()

	// create custom server
	s := &http.Server{
		Addr:           ":8080",
		Handler:        router,
		ReadTimeout:    10 & time.Second,
		WriteTimeout:   10 * time.Second,
		MaxHeaderBytes: 1 << 20,
	}

	// start the HTTP server
	s.ListenAndServe()
}
