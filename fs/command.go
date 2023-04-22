package fs

import (
	"bufio"
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

func FsCommand() *cli.Command {
	return &cli.Command{
		Name:  "fs",
		Usage: "Opens all the currently available accounts and manages them into a filesystem-like style.",
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "one-line",
				Aliases: []string{"ol"},
				Usage:   "Runs a single command and terminates",
			},
		},
		Action: runRepl,
	}
}

func runRepl(ctx *cli.Context) error {
	continueRunning := true
	for continueRunning {
		fmt.Print("> ")
		scanner := bufio.NewScanner(os.Stdin)

		var input string
		if scanner.Scan() {
			input = scanner.Text()
		}

		// Process the input. Parsing the commands.
		tokens, err := Tokenize(input)
		if err != nil {
			return err
		}
		for _, r := range tokens {
			fmt.Println(r)
		}

		if err := processCommand(tokens); err != nil {
			fmt.Println("error: " + err.Error())
		}
	}

	return nil
}
