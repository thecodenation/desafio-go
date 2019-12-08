package main

import (
	"log"
	"net/http"
	"os"
	"encoding/json"
	"math/rand"

	"github.com/codegangsta/negroni"
	"github.com/eucleciojosias/codenation-challenge/pkg/quote"
	"github.com/eucleciojosias/codenation-challenge/pkg/entity"
	"github.com/eucleciojosias/codenation-challenge/pkg/middleware"
	"github.com/gorilla/context"
	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()

	n := negroni.New(
		negroni.HandlerFunc(middleware.Pipeline),
		negroni.NewLogger(),
	)

	quoteRepo := quote.NewSqliteRepository()
	quoteService := quote.NewService(quoteRepo)

	router.Handle("/quote", n.With(
		negroni.Wrap(getRandomQuote(quoteService)),
	)).Methods("GET", "OPTIONS").Name("getRandomQuote")

	router.Handle("/quote/{actor}", n.With(
		negroni.Wrap(getRandomQuote(quoteService)),
	)).Methods("GET", "OPTIONS").Name("getRandomQuote")

	http.Handle("/", router)

	logger := log.New(os.Stderr, "logger: ", log.Lshortfile)
	srv := &http.Server{
		Addr:         ":8080",
		Handler:      context.ClearHandler(http.DefaultServeMux),
		ErrorLog:     logger,
	}

	log.Fatal(srv.ListenAndServe())
}

func getRandomQuote(service quote.Repository) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var allQuotes []*entity.Quote
		var err error

		allQuotes, err = getQuotes(service, r)

		if err != nil && err != entity.NotFound {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error when getting quote"))
			return
		}

		if allQuotes == nil {
			w.WriteHeader(http.StatusNotFound)
			w.Write([]byte("Nothing found"))
			return
		}

		randomID := rand.Int() % len(allQuotes)
		if err := json.NewEncoder(w).Encode(allQuotes[randomID].Detail); err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			w.Write([]byte("Error to return response"))
		}
	})
}

func getQuotes(service quote.Repository, r *http.Request) ([]*entity.Quote, error) {
	params := mux.Vars(r)
	actor, given := params["actor"]
	if given {
		return service.FindByActor(actor)
	}

	return service.FindAll()
}
