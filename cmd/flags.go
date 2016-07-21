package cmd

import (
	"github.com/codegangsta/cli"
	"github.com/webhippie/mygithub/config"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		cli.BoolTFlag{
			Name:        "update, u",
			Usage:       "Enable auto updates",
			EnvVar:      "MYGITHUB_UPDATE",
			Destination: &config.Update,
		},
		cli.StringFlag{
			Name:        "token, t",
			Usage:       "GitHub access token",
			EnvVar:      "MYGITHUB_TOKEN",
			Destination: &config.Token,
			Value:       "",
		},
	}
}
