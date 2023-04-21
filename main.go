package main

import (
	"log"
	"os"

	"ezdrive/auth"
	"ezdrive/config"

	"github.com/urfave/cli/v2"
)

func getCommands() []*cli.Command {
	return []*cli.Command{
		auth.AuthCommand(),
	}
}

func realMain() int {
	if err := config.LoadAppConfig(); err != nil {
		log.Fatalln("error loading config: " + err.Error())
		return 1
	}

	app := &cli.App{
		Name:     "ezdrive",
		Usage:    "Simple CLI tool for managing items in your Google Drive accounts",
		Commands: getCommands(),
	}
	if err := app.Run(os.Args); err != nil {
		log.Fatalln("error: " + err.Error())
		return 1
	}

	return 0
}

func main() {
	os.Exit(realMain())
}
