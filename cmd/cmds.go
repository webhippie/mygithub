package cmd

import (
	"fmt"

	"github.com/codegangsta/cli"
	"github.com/google/go-github/github"
	"github.com/webhippie/mygithub/config"
)

// Commands defines all available sub-commands for this tool.
func Commands() []cli.Command {
	return []cli.Command{
		{
			Name:      "list",
			Aliases:   []string{"l"},
			Usage:     "List available repositories",
			ArgsUsage: "<user/org>",
			Action: func(c *cli.Context) {
				if len(c.Args()) < 1 {
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
					paged, response, err := config.Client.Repositories.List(c.Args().First(), opt)

					if err != nil {
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
}
