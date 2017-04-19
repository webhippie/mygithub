package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/google/go-github/github"
	"github.com/webhippie/mygithub/config"
	"golang.org/x/net/context"
	"golang.org/x/oauth2"
	"gopkg.in/urfave/cli.v2"
)

// Before gets called before any action on every execution.
func Before() cli.BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		if c.Bool("update") {
			Update()
		}

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
