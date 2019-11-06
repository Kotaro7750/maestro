package main

import (
	"maestrocli/command"
	"maestrocli/setting"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "maestrocli"
	app.Usage = "A Command Line Tool for maestro"
	app.Version = "0.1"

	app.Commands = []cli.Command{
		{
			Name:  "init",
			Usage: "initialize server information",
			Action: func(c *cli.Context) error {
				setting.Initial()
				return nil
			},
		},
		{
			Name:  "hello",
			Usage: "hello",
			Action: func(c *cli.Context) error {
				command.Hello()
				return nil
			},
		},
	}

	app.Run(os.Args)
}
