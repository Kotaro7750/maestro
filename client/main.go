package main

import (
	"fmt"
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
				fmt.Printf("aaaa\n")
				return nil
			},
		},
	}

	app.Run(os.Args)
}
