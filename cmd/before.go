package cmd

import (
	"os"

	"github.com/Sirupsen/logrus"
	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/webhippie/mygithub/config"
	"golang.org/x/oauth2"
)

func Before() BeforeFunc {
	return func(c *cli.Context) error {
		logrus.SetOutput(os.Stdout)
		logrus.SetLevel(logrus.DebugLevel)

		if c.BoolT("update") {
			Update()
		}

		if config.Token != "" {
			config.Client = github.NewClient(
				oauth2.NewClient(
					oauth2.NoContext,
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
