package cli

import (
	"fmt"
	"os"

	"github.com/urfave/cli"

	"github.com/ulule/ipfix"
)

// Run runs the application.
func Run() {
	app := cli.NewApp()
	app.Name = "ipfix"
	app.Author = "thoas"
	app.Email = "florent@ulule.com"
	app.Usage = "A webservice to retrieve geolocation information from an ip address"
	app.Version = fmt.Sprintf("%s [git:%s:%s]\ncompiled using %s at %s)", ipfix.Version, ipfix.Branch, ipfix.Revision, ipfix.Compiler, ipfix.BuildTime)
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:   "config, c",
			Usage:  "Config file path",
			EnvVar: "IPFIX_CONF",
		},
	}
	app.Commands = []cli.Command{
		{
			Name:      "version",
			ShortName: "v",
			Usage:     "Retrieve the version number",
			Action: func(c *cli.Context) {
				fmt.Printf("ipfix %s\n", ipfix.Version)
			},
		},
	}
	app.Action = func(c *cli.Context) {
		config := c.String("config")

		if config != "" {
			if _, err := os.Stat(config); err != nil {
				fmt.Fprintf(os.Stderr, "Can't find config file `%s`\n", config)
				os.Exit(1)
			}
		} else {
			fmt.Fprintf(os.Stderr, "Can't find config file\n")
			os.Exit(1)
		}

		err := ipfix.Run(config)

		if err != nil {
			fmt.Fprint(os.Stderr, err)
			os.Exit(1)
		}
	}

	app.Run(os.Args)
}