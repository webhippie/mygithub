package config

import (
	"github.com/google/go-github/github"
)

var (
	Update bool
	Token  string
	Client *github.Client
)
