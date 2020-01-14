package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/AntoineViau/lbcfb/api"
	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
	"github.com/go-chi/render"
)

func routes() *chi.Mux {
	router := chi.NewRouter()
	router.Use(
		render.SetContentType(render.ContentTypeJSON),
		middleware.Logger,
		middleware.DefaultCompress,
		middleware.RedirectSlashes,
		middleware.Recoverer,
	)
	router.Route("/api", func(r chi.Router) {
		r.Mount("/v1", api.FizzBuzzRoutes())
	})
	return router
}

// StartServer allows server launch from another source/package
func StartServer(bindAddress string, port int) {
	router := routes()
	log.Fatal(http.ListenAndServe(fmt.Sprintf("%s:%d", bindAddress, port), router))
}

func main() {
	bindAddress := "0.0.0.0"
	port := 8080
	fmt.Printf("Starting server on %s:%d\n", bindAddress, port)
	StartServer(bindAddress, port)
}
