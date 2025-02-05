package command

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/google/go-github/v69/github"
	"github.com/rs/zerolog/log"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var (
	listCmd = &cobra.Command{
		Use:     "list [name]",
		Aliases: []string{"l"},
		Short:   "List available repositories",
		Run:     listAction,
		Args: func(_ *cobra.Command, args []string) error {
			if len(args) != 1 {
				return fmt.Errorf("missing name argument")
			}

			return nil
		},
	}
)

func init() {
	rootCmd.AddCommand(listCmd)

	listCmd.PersistentFlags().String("format", "", "Output format")
	viper.SetDefault("list.format", "")
	viper.BindPFlag("list.format", listCmd.PersistentFlags().Lookup("format"))
}

func listAction(ccmd *cobra.Command, args []string) {
	client := Client(ccmd.Context())

	opt := &github.RepositoryListByUserOptions{
		ListOptions: github.ListOptions{
			PerPage: 50,
		},
	}

	var repos []*github.Repository

	for {
		paged, response, err := client.Repositories.ListByUser(
			ccmd.Context(),
			args[0],
			opt,
		)

		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to fetch repos")

			os.Exit(1)
		}

		repos = append(repos, paged...)

		if response.NextPage == 0 {
			break
		}

		opt.ListOptions.Page = response.NextPage
	}

	switch format := viper.GetString("list.format"); format {
	case "json":
		val, err := json.MarshalIndent(repos, "", "  ")

		if err != nil {
			log.Error().
				Err(err).
				Msg("failed to format repos")

			os.Exit(1)
		}

		fmt.Println(string(val))
	case "name":
		for _, repo := range repos {
			fmt.Println(*repo.Name)
		}
	default:
		for _, repo := range repos {
			log.Info().
				Msg(*repo.Name)
		}
	}
}
