package github

import (
	"context"

	gh "github.com/google/go-github/github"
)

// Client -- a client structure
type Client struct {
	Client *gh.Client
}

// NewGithubClient -- a client for interacting with GitHub API
func NewGithubClient(ctx context.Context) (Client, error) {
	client := gh.NewClient(nil)
	return Client{
		Client: client,
	}, nil
}
