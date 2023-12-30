package main

import (
	"fmt"
	"github.com/mdobydullah/go-spinner"
	"github.com/urfave/cli/v2"
	"log"
	"os"
	"time"
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

				// Email validation
				spin := spinner.StartNew("Validate: processing...")
				time.Sleep(1 * time.Second)
				valid := validateEmail(email)
				if !valid {
					spin.Stop()
					fmt.Printf("x Validate: %s is an invalid email!", email)
					return nil
				}
				spin.Stop()
				fmt.Println("✓ Validate: completed")

				// Verify email's domain
				verifyDomain(spin, email)

				// Verify is it real or not
				verifyEmail(spin, email)
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
