package cmd

import (
	"github.com/urfave/cli/v2"
)

func init() {

}

var App = &cli.App{
	Name:  "nest",
	Usage: "Nest is a tool for helping develop with go-kit app framework.",
	Action: func(c *cli.Context) error {
		cli.ShowAppHelp(c)
		return nil
	},
}

func AddCommand(cmd *cli.Command) {
	App.Commands = append(App.Commands, cmd)
}
