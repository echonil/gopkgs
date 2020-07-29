package cmd

import (
	"log"
	"os"

	"clevergo.tech/osenv"
	"github.com/echonil/gopkgs/internal/core"
	"github.com/jmoiron/sqlx"
	"github.com/joho/godotenv"
	"github.com/urfave/cli/v2"
)

var (
	db  *sqlx.DB
	app = &cli.App{
		EnableBashCompletion: true,
		Version:              core.Version,
		Name:                 "gopkgs",
		Usage:                "Go Packages",
		Before: func(c *cli.Context) error {
			if err := godotenv.Load(); err != nil {
				return err
			}

			if err := initDB(); err != nil {
				return err
			}

			return nil
		},
	}
)

// Execute executes commands.
func Execute() {
	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}
}

func initDB() (err error) {
	db, err = sqlx.Open("mysql", osenv.Get("MYSQL_DNS"))
	return
}
