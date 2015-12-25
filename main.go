package main

import (
	"fmt"
	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"golang.org/x/oauth2"
	"os"
)

var (
	version  string
	token    string
	username string

	client *github.Client
)

func main() {
	cli.HelpFlag = cli.BoolFlag{
		Name:  "help, h",
		Usage: "Show the help, so what you see now",
	}

	cli.VersionFlag = cli.BoolFlag{
		Name:  "version, v",
		Usage: "Print the current version of that tool",
	}

	app := cli.NewApp()
	app.Name = "mygithub"
	app.Version = version
	app.Author = "Thomas Boerger <thomas@webhippie.de>"
	app.Usage = "Some tiny GitHub client utilities for daily work"

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:        "token, t",
			Usage:       "GitHub access token",
			EnvVar:      "MYGITHUB_TOKEN",
			Destination: &token,
			Value:       "",
		},
	}

	app.Before = func(context *cli.Context) error {
		if token != "" {
			ts := oauth2.StaticTokenSource(
				&oauth2.Token{
					AccessToken: token,
				},
			)

			tc := oauth2.NewClient(
				oauth2.NoContext,
				ts,
			)

			client = github.NewClient(tc)
		} else {
			client = github.NewClient(nil)
		}

		return nil
	}

	app.Commands = []cli.Command{
		{
			Name:    "list",
			Aliases: []string{"l"},
			Usage:   "List available repositories",
			Action: func(c *cli.Context) {
				if len(c.Args()) > 0 {
					username = c.Args().First()
				} else {
					fmt.Println("Please provide a username or org")
					return
				}

				opt := &github.RepositoryListOptions{
					ListOptions: github.ListOptions{
						PerPage: 50,
					},
				}

				var repos []github.Repository

				for {
					paged, response, err := client.Repositories.List(username, opt)

					if err != nil {
						fmt.Println(response)
						fmt.Println(err)
						return
					}

					repos = append(repos, paged...)

					if response.NextPage == 0 {
						break
					}

					opt.ListOptions.Page = response.NextPage
				}

				for _, repo := range repos {
					fmt.Println(*repo.Name)
				}
			},
		},
	}

	app.Run(os.Args)
}
