package config

import (
	"github.com/google/go-github/github"
	"golang.org/x/net/context"
)

var (
	// Update determines if automated updates are enabled.
	Update bool

	// Token defines the GitHub API token to avoid API blocks.
	Token string

	// Client represents the current GitHub API client.
	Client *github.Client

	// Context is used for the GitHub API request, required by the client.
	Context context.Context
)
