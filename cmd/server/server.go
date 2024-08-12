package main

import (
	"net/http"

	"github.com/joncalhoun/global-refactor-demo/psql"
	"github.com/joncalhoun/global-refactor-demo/tweet"
)

func main() {
	tweetStorage := tweet.TweetStorage{DB: psql.DB}
	mux := http.NewServeMux()
	mux.HandleFunc("GET /tweets/new", tweet.NewTweetHandler())
	mux.HandleFunc("GET /tweets/{id}", tweet.ShowTweetHandler(tweetStorage))
	mux.HandleFunc("POST /tweets", tweet.CreateTweetHandler(tweetStorage))
	http.ListenAndServe(":8080", mux)
}
