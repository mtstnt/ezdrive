package auth

import (
	"ezdrive/config"
	"ezdrive/utils"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/urfave/cli/v2"
	"golang.org/x/oauth2/google"
	"google.golang.org/api/drive/v3"
)

func AuthCommand() *cli.Command {
	return &cli.Command{
		Name:  "auth",
		Usage: "Subset of commands that deal with account management.",
		Subcommands: []*cli.Command{
			{
				Name:   "login",
				Usage:  "Log in with an account, labeling it with a nickname for local use.",
				Action: loginCommand,
			},
			{
				Name:   "logout",
				Usage:  "Log out an account given its specific nickname.",
				Action: logoutCommand,
			},
		},
	}
}

// [Invokes command `ezdrive auth login [account_nickname]`]
//
// Requires:
//
// - Nickname. Must be unique for all accounts. How? Check if the filename is unique in cached/creds/ folder.
//
// Stdin:
//
// - Path to credentials.json for the account. Will be cached locally in `APP_ROOT/cached/creds/[nickname].json`
// - Authentication token from URL. Will be cached locally in `APP_ROOT/cached/tokens/[nickname].json`
func loginCommand(ctx *cli.Context) error {
	// Required argument.
	var (
		nickname        string
		credentialsPath string = ""
	)

	nickname = ctx.Args().First()
	if nickname == "" {
		return fmt.Errorf("nickname must be supplied")
	}

	if ctx.Args().Len() >= 2 {
		credentialsPath = ctx.Args().Get(1)
	}

	// Check if nickname is unique.
	isUnique, err := isNicknameUnique(nickname, config.CredsDir)
	if err != nil {
		return err
	}
	if !isUnique {
		return fmt.Errorf("nickname is not unique")
	}

	// Ask for path to the credentials.json.
	if credentialsPath == "" {
		fmt.Print("Paste your credentials.json path here: ")
		if _, err := fmt.Scanf("%s\n", &credentialsPath); err != nil {
			return err
		}
	}

	// Retrieve config from credentials.json path and store it a copy.
	newCredsPath := filepath.Join(config.CredsDir, nickname+".json")
	newTokenPath := filepath.Join(config.TokensDir, nickname+".json")

	if err := utils.CopyFile(credentialsPath, newCredsPath); err != nil {
		return err
	}

	// creds/nickname.json
	fptr, err := os.Open(newCredsPath)
	if err != nil {
		return err
	}
	content, err := io.ReadAll(fptr)
	if err != nil {
		return err
	}

	config, err := google.ConfigFromJSON(content, drive.DriveScope)
	if err != nil {
		return err
	}
	token, err := promptTokenInWeb(config)
	if err != nil {
		return err
	}

	saveToken(newTokenPath, token)

	return nil
}

// Invokes command `ezdrive auth logout [account_nickname|* (to delete all)]`
func logoutCommand(*cli.Context) error {
	return nil
}
