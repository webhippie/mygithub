package main

import (
	"github.com/webhippie/mygithub/pkg/config"
	"gopkg.in/urfave/cli.v2"
)

// Flags defines all available flags for this command.
func Flags() []cli.Flag {
	return []cli.Flag{
		&cli.StringFlag{
			Name:        "token, t",
			Usage:       "GitHub access token",
			EnvVars:     []string{"MYGITHUB_TOKEN"},
			Destination: &config.Token,
			Value:       "",
		},
	}
}
