package main

import (
	"context"
	"os"

	"github.com/google/go-github/github"
	"github.com/sirupsen/logrus"
	"github.com/webhippie/mygithub/pkg/config"
	"golang.org/x/oauth2"
	"gopkg.in/urfave/cli.v2"
)

// Before gets called before any action on every execution.
func Before() cli.BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		config.Context = context.Background()

		if config.Token != "" {
			config.Client = github.NewClient(
				oauth2.NewClient(
					config.Context,
					oauth2.StaticTokenSource(
						&oauth2.Token{
							AccessToken: config.Token,
						},
					),
				),
			)
		} else {
			config.Client = github.NewClient(nil)
		}

		return nil
	}
}
