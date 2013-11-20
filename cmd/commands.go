package cmd

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:      "ls-remote",
		ShortName: "lsr",
		Usage:     "List all available Note.js versions",
		Action:    lsRemote,
	},
	{
		Name:      "install",
		ShortName: "i",
		Usage:     "Install specific Note.js version",
		Action:    install,
	},
	{
		Name:   "ls",
		Usage:  "List all installed Node.js versions",
		Action: lsLocal,
	},
}
