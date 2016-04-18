package cmd

import (
	"fmt"
	"os"

	"github.com/sanbornm/go-selfupdate/selfupdate"
	"github.com/webhippie/mygithub/config"
)

var (
	updates string = "http://dl.webhippie.de/"
)

func Update() {
	if config.VersionDev == "dev" {
		fmt.Fprintf(os.Stderr, "Updates are disabled for development versions.\n")
	} else {
		updater := &selfupdate.Updater{
			CurrentVersion: fmt.Sprintf(
				"%d.%d.%d",
				config.VersionMajor,
				config.VersionMinor,
				config.VersionPatch,
			),
			ApiURL:  updates,
			BinURL:  updates,
			DiffURL: updates,
			Dir:     "updates/",
			CmdName: "mygithub",
		}

		go updater.BackgroundRun()
	}
}
