package repo

import (
	esv8 "github.com/elastic/go-elasticsearch/v8"
)

type MovieRepo struct {
	client *esv8.Client
	index  string
}

func NewMovieEsRepo(client *esv8.Client) *MovieRepo {
	return &MovieRepo{
		client: client,
		index:  "movieswiki",
	}
}
