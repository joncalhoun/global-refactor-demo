package main

import (
	"net/http"

	"github.com/joncalhoun/global-refactor-demo/psql"
	"github.com/joncalhoun/global-refactor-demo/tweet"
)

func main() {
	mux := http.NewServeMux()
	// mux.HandleFunc("GET /tweets/{id}", tweet.ShowTweetHandler())
	// mux.HandleFunc("GET /tweets/new", tweet.NewTweetHandler())
	mux.HandleFunc("POST /tweets", tweet.CreateTweetHandler(psql.DB))
	http.ListenAndServe(":8080", mux)
}
