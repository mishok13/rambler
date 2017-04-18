package main

import (
	"fmt"
	"os"

	_ "github.com/sbp-contrib/rambler/driver/mysql"
	_ "github.com/sbp-contrib/rambler/driver/postgresql"
	_ "github.com/sbp-contrib/rambler/driver/sqlite"
	"github.com/sbp-contrib/rambler/rambler"
	"github.com/urfave/cli"
)

var app *cli.App

// VERSION holds the version of rambler as defined at compile time.
var VERSION string

func main() {
	var app = cli.NewApp()

	app.Name = "rambler"
	app.Usage = "Migrate all the things!"
	app.Version = VERSION
	app.Authors = []cli.Author{
		{
			Name:  "Romain Baugue",
			Email: "romain.baugue@elwinar.com",
		},
	}

	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "configuration, c",
			Value: "rambler.json",
			Usage: "path to the configuration file",
		},
		cli.StringFlag{
			Name:  "environment, e",
			Value: "default",
			Usage: "set the working environment",
		},
		cli.BoolFlag{
			Name:  "debug",
			Usage: "display debug messages",
		},
	}
	app.Before = rambler.Bootstrap
	app.Commands = []cli.Command{
		{
			Name:  "apply",
			Usage: "apply the next migration",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Apply all migrations",
				},
			},
			Action: rambler.Apply,
		},
		{
			Name:  "reverse",
			Usage: "reverse the last migration",
			Flags: []cli.Flag{
				cli.BoolFlag{
					Name:  "all, a",
					Usage: "Reverse all migrations",
				},
			},
			Action: rambler.Reverse,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
