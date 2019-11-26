package cmd

import (
	"fmt"

	"github.com/urfave/cli/v2"
)

func init() {
	AddCommand(initCmd)
}

var initCmd = &cli.Command{
	Name:      "init",
	Usage:     "Initialize a go-kit application",
	HelpName:  "init",
	UsageText: "nest init [name] [flags]",
	Flags:     []cli.Flag{
		//		&cli.BoolFlag{Name: "", Aliases: []string{"forevvarr"}},
	},
	SkipFlagParsing: false,
	HideHelp:        false,
	Hidden:          false,
	Before: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "brace for impact\n")
		return nil
	},
	After: func(c *cli.Context) error {
		fmt.Fprintf(c.App.Writer, "did we lose anyone?\n")
		return nil
	},
	Action: func(c *cli.Context) error {
		c.Command.FullName()
		c.Command.HasName("wop")
		c.Command.Names()
		c.Command.VisibleFlags()
		fmt.Fprintf(c.App.Writer, "ddodododooo\n")
		if c.Bool("forever") {
			c.Command.Run(c)
		}
		return nil
	},
	OnUsageError: func(c *cli.Context, err error, isSubcommand bool) error {
		fmt.Fprintf(c.App.Writer, "for shame\n")
		return err
	},
}
