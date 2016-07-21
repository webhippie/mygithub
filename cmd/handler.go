package cmd

import (
	"github.com/codegangsta/cli"
)

// BeforeFunc is a general callback function for the before handling.
type BeforeFunc func(*cli.Context) error
