package tweet

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func NewTweetHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprint(w, `
		<html>
		<body>
		<form action="/tweets" method="POST">
  		<textarea name="content"></textarea>
  		<button type="submit">Submit</button>
		</form>
		</body>
		</html>`)
	}
}

func ShowTweetHandler(storage TweetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		idStr := r.PathValue("id")
		id, err := strconv.Atoi(idStr)
		if err != nil {
			http.Error(w, "invalid tweet id", http.StatusBadRequest)
			return
		}
		tweet, err := storage.GetTweet(id)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tweet)
	}
}

func CreateTweetHandler(storage TweetStorage) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		r.ParseForm()
		creatorID := 1 // In a real app we would get this from the session
		content := r.Form.Get("content")
		tweet, err := storage.CreateTweet(creatorID, content)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}
		json.NewEncoder(w).Encode(tweet)
	}
}
