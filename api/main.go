package main

import (
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"os"

	"github.com/codegangsta/negroni"
	"github.com/eucleciojosias/codenation-challenge/pkg/entity"
	"github.com/eucleciojosias/codenation-challenge/pkg/middleware"
	"github.com/eucleciojosias/codenation-challenge/pkg/quote"
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
		Addr:     ":8080",
		Handler:  context.ClearHandler(http.DefaultServeMux),
		ErrorLog: logger,
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
		singleQuote := allQuotes[randomID]

		response := make(map[string]string)
		response["actor"] = singleQuote.Actor
		response["quote"] = singleQuote.Detail

		if err := json.NewEncoder(w).Encode(response); err != nil {
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
