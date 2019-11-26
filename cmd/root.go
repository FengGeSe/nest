package cmd

import (
	"fmt"
	"os"

	"github.com/urfave/cli/v2"
)

var rootApp = &cli.App{
	Name:    "nest",
	Version: "v1.0.0",
	Usage:   "Nest is a tool for helping develop with go-kit app framework.",
	Action: func(c *cli.Context) error {
		fmt.Println("boom! i say!")
		return nil
	},
}

func AddCommand(cmd *cli.Command) {
	rootApp.Commands = append(rootApp.Commands, cmd)
}

func Run() error {
	return rootApp.Run(os.Args)
}
