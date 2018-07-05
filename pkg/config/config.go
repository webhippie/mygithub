package config

import (
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

var (
	// Token defines the GitHub API token to avoid API blocks.
	Token string

	// Client represents the current GitHub API client.
	Client *github.Client

	// Context is used for the GitHub API request, required by the client.
	Context context.Context
)
