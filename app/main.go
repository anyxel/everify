package main

import (
	"fmt"
	"github.com/urfave/cli/v2"
	"log"
	"os"
)

var (
	version string
	build   string
)

func main() {
	app := &cli.App{
		Name:    "everify",
		Version: version,
		Usage:   "Anyxel Email Verify",

		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "email",
				Aliases: []string{"e"},
				Usage:   "verify email address",
			},
		},
		Action: func(cCtx *cli.Context) error {
			if len(cCtx.String("email")) != 0 {
				email := cCtx.String("email")

				run(email)
			} else {
				fmt.Println("Anyxel Email Verify.")
				fmt.Println("Version: ", version)
				fmt.Println("Build: ", build)
			}

			return nil
		},
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
