package cmd

import (
	"fmt"

	"github.com/google/go-github/github"
	"github.com/webhippie/mygithub/config"
	"gopkg.in/urfave/cli.v2"
)

// Commands defines all available sub-commands for this tool.
func Commands() []*cli.Command {
	return []*cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "List available repositories",
			ArgsUsage: "<user/org>",
			Action: func(c *cli.Context) error {
				if c.Args().Len() < 1 {
					return fmt.Errorf("Please provide a username or org")
				}

				opt := &github.RepositoryListOptions{
					ListOptions: github.ListOptions{
						PerPage: 50,
					},
				}

				var repos []*github.Repository

				for {
					paged, response, err := config.Client.Repositories.List(config.Context, c.Args().First(), opt)

					if err != nil {
						return err
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

				return nil
			},
		},
	}
}
