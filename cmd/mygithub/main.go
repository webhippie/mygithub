package main

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/webhippie/mygithub/pkg/command"
)

func main() {
	if env := os.Getenv("MYGITHUB_ENV_FILE"); env != "" {
		godotenv.Load(env)
	}

	if err := command.Run(); err != nil {
		os.Exit(1)
	}
}
