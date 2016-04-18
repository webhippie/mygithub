package cmd

import (
	"github.com/codegangsta/cli"
)

type BeforeFunc func(*cli.Context) error
