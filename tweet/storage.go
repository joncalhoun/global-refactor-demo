package tweet

import (
	"database/sql"
	"fmt"
)

type Tweet struct {
	ID        int
	CreatorID int
	Content   string
}

type TweetStorage struct {
	DB *sql.DB
}

func (ts TweetStorage) CreateTweet(creatorID int, content string) (*Tweet, error) {
	row := ts.DB.QueryRow(`
		INSERT INTO tweets (creator_id, content)
		VALUES ($1, $2)
		RETURNING id`, creatorID, content)
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to insert tweet into database: %v", err)
	}
	tweet := Tweet{
		CreatorID: creatorID,
		Content:   content,
	}
	err = row.Scan(&tweet.ID)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to scan id from database: %v", err)
	}
	return &tweet, nil
}

func (ts TweetStorage) GetTweet(id int) (*Tweet, error) {
	row := ts.DB.QueryRow(`
		SELECT creator_id, content
		FROM tweets
		WHERE id = $1`, id)
	err := row.Err()
	if err != nil {
		return nil, fmt.Errorf(
			"failed to query database: %v", err)
	}
	tweet := Tweet{ID: id}
	err = row.Scan(&tweet.CreatorID, &tweet.Content)
	if err != nil {
		return nil, fmt.Errorf(
			"failed to scan tweet from database: %v", err)
	}
	return &tweet, nil
}
