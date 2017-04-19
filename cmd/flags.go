package cmd

import (
	"github.com/webhippie/mygithub/config"
	"gopkg.in/urfave/cli.v2"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.BoolFlag{
			Name:        "update, u",
			Value:       true,
			Usage:       "Enable auto updates",
			EnvVars:     []string{"MYGITHUB_UPDATE"},
			Destination: &config.Update,
		},
		&cli.StringFlag{
			Name:        "token, t",
			Usage:       "GitHub access token",
			EnvVars:     []string{"MYGITHUB_TOKEN"},
			Destination: &config.Token,
			Value:       "",
		},
	}
}
