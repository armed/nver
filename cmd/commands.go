package cmd

import "github.com/codegangsta/cli"

var Commands = []cli.Command{
	{
		Name:      "install",
		ShortName: "i",
		Usage:     "Install specific Note.js version",
		Action:    Install,
	},
	{
		Name:   "use",
		Usage:  "Create symlink for specific Note.js version",
		Action: Use,
	},
	{
		Name:      "remove",
		ShortName: "rm",
		Usage:     "Remove installed Node.js version",
		Action:    Remove,
	},
	{
		Name:      "ls-remote",
		ShortName: "lsr",
		Usage:     "List all available Note.js versions",
		Action:    LsRemote,
	},
	{
		Name:   "ls",
		Usage:  "List all installed Node.js versions",
		Action: LsLocal,
	},
}
