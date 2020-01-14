package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/AntoineViau/lbcfb/fizzbuzz"
	"github.com/go-chi/chi"
	"github.com/go-chi/render"
)

// FizzBuzzRoutes will provide the Chi routes for the server.
func FizzBuzzRoutes() *chi.Mux {
	router := chi.NewRouter()
	router.Get("/fizzbuzz", fizzbuzzHandler)
	return router
}

func fizzbuzzHandler(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query()
	for _, paramName := range []string{"string1", "string2", "int1", "int2", "limit"} {
		if query[paramName] == nil {
			http.Error(w, fmt.Sprintf("Missing parameter %s", paramName), 400)
			return
		}
	}
	string1 := query.Get("string1")
	string2 := query.Get("string2")
	int1, err := strconv.Atoi(query.Get("int1"))
	if err != nil {
		http.Error(w, "Invalid parameter int1", 400)
		return
	}
	if int1 == 0 {
		http.Error(w, "int1 can't be 0", 400)
		return
	}
	int2, err := strconv.Atoi(query.Get("int2"))
	if err != nil {
		http.Error(w, "Invalid parameter int2", 400)
		return
	}
	if int2 == 0 {
		http.Error(w, "int2 can't be 0", 400)
		return
	}
	limit, err := strconv.Atoi(query.Get("limit"))
	if err != nil || limit < 0 {
		http.Error(w, "Invalid parameter limit", 400)
		return
	}
	output, err := fizzbuzz.FizzBuzz(string1, string2, int1, int2, limit)
	if err != nil {
		http.Error(w, http.StatusText(500), 500)
		return
	}
	render.JSON(w, r, output)
}
